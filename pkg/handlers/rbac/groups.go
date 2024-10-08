package rbac

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GroupsHandler handles requests related to listing groups.
func GroupsHandler(clientset *kubernetes.Clientset) echo.HandlerFunc {
	return func(c echo.Context) error {
		roleBindings, err := clientset.RbacV1().RoleBindings("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error listing role bindings: "+err.Error())
		}

		clusterRoleBindings, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error listing cluster role bindings: "+err.Error())
		}

		groups := extractGroupsFromBindings(roleBindings.Items, clusterRoleBindings.Items)
		return c.JSON(http.StatusOK, groups)
	}
}

// extractGroupsFromBindings extracts groups from RoleBindings and ClusterRoleBindings.
func extractGroupsFromBindings(roleBindings []rbacv1.RoleBinding, clusterRoleBindings []rbacv1.ClusterRoleBinding) []string {
	groupSet := make(map[string]struct{})

	for _, rb := range roleBindings {
		for _, subject := range rb.Subjects {
			if subject.Kind == rbacv1.GroupKind {
				groupSet[subject.Name] = struct{}{}
			}
		}
	}

	for _, crb := range clusterRoleBindings {
		for _, subject := range crb.Subjects {
			if subject.Kind == rbacv1.GroupKind {
				groupSet[subject.Name] = struct{}{}
			}
		}
	}

	groups := make([]string, 0, len(groupSet))
	for group := range groupSet {
		groups = append(groups, group)
	}

	return groups
}
