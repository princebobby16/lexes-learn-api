package router

import (
	"lexes_learn_server/app/controllers/creaters"
	"lexes_learn_server/app/controllers/getters"
	"lexes_learn_server/app/controllers/healthcheck"
	"lexes_learn_server/app/controllers/login"
	"lexes_learn_server/app/controllers/setters"
	"net/http"
)

//Route Create a single route object
type Route struct {
	Name string
	Path string
	Method string
	Handler http.HandlerFunc
}

//Routes Create an object of different routes
type Routes []Route

// InitRoutes Set up routess
func InitRoutes() Routes {
	routes := Routes{
		// health check
		Route{
			Name:    "Index",
			Path:    "/",
			Method:  http.MethodGet,
			Handler: healthcheck.IndexHandler,
		},
		// login
		Route{
			Name:    "LoginStudent",
			Path:    "/users/student/login",
			Method:  http.MethodPost,
			Handler: login.SignInStudentHandler,
		},
		Route{
			Name:    "LoginTeacher",
			Path:    "/users/teacher/login",
			Method:  http.MethodPost,
			Handler: login.SignInTeacherHandler,
		},
		// announcement
		Route{
			Name:    "Get Announcements",
			Path:    "/get-announcement",
			Method:  http.MethodPost,
			Handler: getters.GetAnnouncements,
		},
		// assignment
		Route{
			Name: "Get all Assignment",
			Path: "/get-all-assignments",
			Method: http.MethodGet,
			Handler: getters.GetAllAssignments,
		},
		Route{
			Name:    "Get one Assignment",
			Path:    "/get-one-assigment",
			Method:  http.MethodGet,
			Handler: getters.GetOneAssignmentById,
		},
		Route{
			Name: "Create assignment",
			Path: "/create-assignment",
			Method: http.MethodPost,
			Handler: creaters.CreateAssignment,
		},
		// question
		Route{
			Name: "Set question",
			Path: "/set-question",
			Method: http.MethodPost,
			Handler: setters.SetQuestion,
		},
		Route{
			Name:    "Get one Question",
			Path:    "/get-one-question/{id}",
			Method:  http.MethodGet,
			Handler: getters.GetOneQuestionById,
		},
		Route{
			Name:    "Get All Question",
			Path:    "/get-all-questions",
			Method:  http.MethodGet,
			Handler: getters.GetAllQuestions,
		},
		// course
		Route{
			Name: "Create course",
			Path: "/create-course",
			Method: http.MethodPost,
			Handler: creaters.CreateCourse,
		},
		Route{
			Name:    "Get one Course",
			Path:    "/get-one-course/{id}",
			Method:  http.MethodGet,
			Handler: getters.GetOneCourseById,
		},
		Route{
			Name:    "Get All Courses",
			Path:    "/get-all-courses",
			Method:  http.MethodGet,
			Handler: getters.GetAllCourses,
		},
		// class
		Route{
			Name:    "Create class",
			Path:    "/create-class",
			Method:  http.MethodPost,
			Handler: creaters.CreateClass,
		},
		Route{
			Name:    "Get one Class",
			Path:    "/get-one-class/{id}",
			Method:  http.MethodGet,
			Handler: getters.GetOneClassById,
		},
		Route{
			Name:    "Get All Class",
			Path:    "/get-all-classes",
			Method:  http.MethodGet,
			Handler: getters.GetAllClasses,
		},
	}

	return routes
}