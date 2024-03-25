package main

import (
	"sort"

	"github.com/fast-admin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MenuNode struct {
	models.SysMenu
	Children []*MenuNode
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/fast_admin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("数据库错误")
	}

	// db.AutoMigrate(
	// 	&models.SysUser{},
	// 	&models.SysRole{},
	// 	&models.SysUserRole{},
	// 	&models.SysMenu{},
	// 	&models.SysRoleMenu{},
	// 	&models.SysDictType{},
	// 	&models.SysDictData{},
	// 	&models.SysLoginLog{},
	// )

	r := gin.Default()

	r.GET("", func(ctx *gin.Context) {
		var menus []models.SysMenu
		err := db.Find(&menus).Error
		if err != nil {
			ctx.JSON(200, err)
			return
		}

		menuTree := GenerateMenuTree(menus)
		ctx.JSON(200, menuTree)
	})

	r.Run(":3000")

}

func GenerateMenuTree(menus []models.SysMenu) []*MenuNode {
	// 构建菜单映射表
	menuMap := make(map[uint]*MenuNode)
	for _, menu := range menus {
		menuMap[menu.ID] = &MenuNode{
			SysMenu:  menu,
			Children: []*MenuNode{},
		}
	}

	// 构建菜单树
	var rootNodes []*MenuNode
	for _, menu := range menuMap {
		if menu.ParentID == 0 {
			rootNodes = append(rootNodes, menu)
		} else {
			parent, ok := menuMap[menu.ParentID]
			if ok {
				parent.Children = append(parent.Children, menu)
			}
		}
	}

	// 对菜单进行排序
	sortMenuNodes(rootNodes)

	return rootNodes
}

func sortMenuNodes(nodes []*MenuNode) {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].OrderNum < nodes[j].OrderNum
	})

	for _, node := range nodes {
		sortMenuNodes(node.Children)
	}
}
