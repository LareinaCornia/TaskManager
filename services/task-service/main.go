package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Task struct {
	ID             int       `db:"id" json:"id"`
	Title          string    `db:"title" json:"title"`
	Description    string    `db:"description" json:"description"`
	Status         string    `db:"status" json:"status"`
	AssignedUserID int       `db:"assigned_user_id" json:"assigned_user_id"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

var (
	db        *sqlx.DB
	jwtSecret = []byte("your_secret_key") // 生产环境不要硬编码
)

func main() {
	var err error
	db, err = sqlx.Connect("mysql", "root:20030808@tcp(127.0.0.1:3306)/taskmanager?parseTime=true")
	if err != nil {
		log.Fatalln("DB connection failed:", err)
	}

	router := gin.Default()
	router.Use(AuthMiddleware())

	taskRoutes := router.Group("/tasks")
	{
		taskRoutes.POST("", CreateTask)
		taskRoutes.PUT("/:id", UpdateTask)
		taskRoutes.GET("", GetTasks)
	}

	router.Run(":8081")
}

func CreateTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.NamedExec(`INSERT INTO tasks (title, description, status, assigned_user_id)
	VALUES (:title, :description, :status, :assigned_user_id)`, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task created"})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task.ID, _ = strconv.Atoi(id)
	_, err := db.NamedExec(`UPDATE tasks SET status=:status WHERE id=:id`, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}

func GetTasks(c *gin.Context) {
	userID := c.Query("user_id")
	var tasks []Task
	err := db.Select(&tasks, "SELECT * FROM tasks WHERE assigned_user_id = ?", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			c.Abort()
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		userID := int((*claims)["user_id"].(float64))
		c.Set("user_id", userID)
		c.Next()
	}
}
