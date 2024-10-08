package rbac

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GroupDetailsResponse represents the detailed information about a group.
type GroupDetailsResponse struct {
	GroupName           string                      `json:"groupName"`
	RoleBindings        []rbacv1.RoleBinding        `json:"roleBindings"`
	ClusterRoleBindings []rbacv1.ClusterRoleBinding `json:"clusterRoleBindings"`
	ClusterRoles        []rbacv1.ClusterRole        `json:"clusterRoles"`
}

// GroupDetailsHandler handles requests for detailed information about a specific group.
func GroupDetailsHandler(clientset *kubernetes.Clientset) echo.HandlerFunc {
	return func(c echo.Context) error {
		groupName := c.QueryParam("groupName")
		if groupName == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "Group name is required")
		}

		roleBindings, err := clientset.RbacV1().RoleBindings("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error listing role bindings: "+err.Error())
		}

		clusterRoleBindings, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error listing cluster role bindings: "+err.Error())
		}

		clusterRoles, err := clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error listing cluster roles: "+err.Error())
		}

		groupDetails := extractGroupDetails(groupName, roleBindings.Items, clusterRoleBindings.Items, clusterRoles.Items)
		return c.JSON(http.StatusOK, groupDetails)
	}
}

// extractGroupDetails extracts detailed information about a specific group.
func extractGroupDetails(groupName string, roleBindings []rbacv1.RoleBinding, clusterRoleBindings []rbacv1.ClusterRoleBinding, clusterRoles []rbacv1.ClusterRole) GroupDetailsResponse {
	var groupRoleBindings []rbacv1.RoleBinding
	var groupClusterRoleBindings []rbacv1.ClusterRoleBinding
	var groupClusterRoles []rbacv1.ClusterRole

	for _, rb := range roleBindings {
		for _, subject := range rb.Subjects {
			if subject.Kind == rbacv1.GroupKind && subject.Name == groupName {
				groupRoleBindings = append(groupRoleBindings, rb)
			}
		}
	}

	for _, crb := range clusterRoleBindings {
		for _, subject := range crb.Subjects {
			if subject.Kind == rbacv1.GroupKind && subject.Name == groupName {
				groupClusterRoleBindings = append(groupClusterRoleBindings, crb)
			}
		}
	}

	// Collect ClusterRoles associated with the group's ClusterRoleBindings
	for _, crb := range groupClusterRoleBindings {
		for _, cr := range clusterRoles {
			if cr.Name == crb.RoleRef.Name {
				groupClusterRoles = append(groupClusterRoles, cr)
			}
		}
	}

	return GroupDetailsResponse{
		GroupName:           groupName,
		RoleBindings:        groupRoleBindings,
		ClusterRoleBindings: groupClusterRoleBindings,
		ClusterRoles:        groupClusterRoles,
	}
}