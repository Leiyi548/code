package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Desc    string `json:"description"`
}

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Age      int    `json:"age" form:"age"`
}

func print(s1, s2 string) string {
	fmt.Print(s1, s2)
	return s1 + "---" + s2
}

func main() {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// r.LoadHTMLGlob("template/*")
	r.LoadHTMLGlob("templates/**/*")

	// 配置静态服务
	r.Static("static", "./static")

	// 自定义模板函数 注意把这个全局函数加在加载模板前
	r.SetFuncMap(template.FuncMap{
		"print": print,
	})

	r.GET("/ping", func(c *gin.Context) {
		// 有json，string就不会显示了！
		c.String(200, "hello,%v", "你好")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		c.String(200, "hello,%v", "你好")
	})

	r.POST("/add", func(c *gin.Context) {
		c.String(200, "这是一个post---请求")
	})

	r.PUT("/edit", func(c *gin.Context) {
		c.String(200, "这是一个put---请求")
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.String(200, "这是一个delete---请求")
	})

	r.GET("/customJson", func(c *gin.Context) {
		customJsonText := &Article{
			Title:   "维尼熊的日记",
			Desc:    "一个想当皇帝的熊",
			Content: "我是皇帝，我是一尊！",
		}
		c.JSON(200, customJsonText)
	})

	r.GET("/customJsonp", func(c *gin.Context) {
		customJsonText := &Article{
			Title:   "维尼熊的日记",
			Desc:    "一个想当皇帝的熊",
			Content: "我是皇帝，我是一尊！ppppppp！",
		}
		c.JSONP(200, customJsonText)
	})

	r.GET("/admin/index", func(c *gin.Context) {
		// r.LoadHTMLGlob("source/*")

		type CCP struct {
			Emperor string `json:"emperor"`
			Year    int32  `json:"year"`
			Desc    string `json:"description"`
		}

		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "admin/index",
			"ccp": &CCP{
				Emperor: "维尼熊",
				Year:    73,
				Desc:    "一个邪恶的组织，剥削压迫，所有东西都是它的，不是你的！",
			},
		})
	})

	r.GET("/default/index", func(c *gin.Context) {

		type ECUT struct {
			Emperor string `json:"emperor"`
			Year    int32  `json:"year"`
			Desc    string `json:"description"`
		}

		c.HTML(http.StatusOK, "default/default.html", gin.H{
			"title": "default/index",
			"ecut": &ECUT{
				Emperor: "可恶的人",
				Year:    63,
				Desc:    "天天封，封尼妈逼啊！！！",
			},
		})
	})

	r.GET("/default/score", func(c *gin.Context) {

		type Student struct {
			Name  string `json:"studentName"`
			Year  int32  `json:"year"`
			Desc  string `json:"description"`
			Score int32  `json:"score"`
			Todo  []string
		}

		c.HTML(http.StatusOK, "default/score.html", gin.H{
			"title": "default/score",
			"student": &Student{
				Name:  "小明",
				Year:  18,
				Score: 59,
				Desc:  "填鸦式教育害死人！！！",
				Todo:  []string{"吃饭", "写作业", "玩游戏"},
			},
		})
	})

	// Get 请求传值
	r.GET("/", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		page := c.DefaultQuery("page", "1")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})

	r.GET("/default/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/user.html", gin.H{})
	})

	r.GET("/getUser", func(c *gin.Context) {
		user := &UserInfo{}
		if err := c.ShouldBind(&user); err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})

	// POST
	r.POST("/doAddUser", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	r.GET("/doAddUser2", func(c *gin.Context) {
		user := &UserInfo{}
		if err := c.ShouldBind(&user); err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})


	// 默认是在8080
	// r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	r.Run(":8000") // 监听并在 0.0.0.0:8080 上启动服务
}
