// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"net/http"

	"github.com/svanellewee/todoserver/models"
	"github.com/svanellewee/todoserver/restapi/operations"
)

//go:generate swagger generate server --target ../../todoserver --name TaskListServer --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.TaskListServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

var tasks = make([]*models.Task, 0, 100)

func addTask(task *models.Task) (*models.Task, error) {
	task.ID = int64(len(tasks)) + 1
	tasks = append(tasks, task)
	return task, nil
}

func getTask(id int64) *models.Task {
	if len(tasks) > int(id)-1 {
		return tasks[id-1]
	}
	return nil
}

func updateTask(id int64, newTask *models.Task) (*models.Task, error) {
	if len(tasks) > int(id)-1 {
		tasks[id-1].Completed = newTask.Completed
		tasks[id-1].Description = newTask.Description
		return tasks[id-1], nil
	}
	return nil, fmt.Errorf("did not find task ID: %d", id)
}
func configureAPI(api *operations.TaskListServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()
	api.UpdateOneHandler = operations.UpdateOneHandlerFunc(func(params operations.UpdateOneParams) middleware.Responder {
		fmt.Printf(">......%v\n", params.ID)

		result, err := updateTask(params.ID, params.Body)
		if err != nil {
			return operations.NewUpdateOneDefault(404).WithPayload(&models.Error{
				Code:    404,
				Message: swag.String(err.Error()), // swag.String(fmt.Sprintf("could not update nonexistent ID %d", params.ID)),
			})
		}
		return operations.NewUpdateOneOK().WithPayload(result)
	})

	api.GetTaskHandler = operations.GetTaskHandlerFunc(func(params operations.GetTaskParams) middleware.Responder {
		return operations.NewGetTaskOK().WithPayload(tasks) //middleware.NotImplemented("operation operations.GetTask has not yet been implemented")
	})
	api.GetTaskIDHandler = operations.GetTaskIDHandlerFunc(func(params operations.GetTaskIDParams) middleware.Responder {
		result := getTask(params.ID)
		if result == nil {
			return operations.NewPostTaskDefault(404).WithPayload(&models.Error{
				Code:    404,
				Message: swag.String(fmt.Sprintf("could not find %d", params.ID)),
			})
		}
		return operations.NewGetTaskIDOK().WithPayload(getTask(params.ID))
	})
	api.PostTaskHandler = operations.PostTaskHandlerFunc(func(params operations.PostTaskParams) middleware.Responder {
		result, err := addTask(params.Body)
		if err != nil {
			return operations.NewPostTaskDefault(500).WithPayload(&models.Error{
				Code:    500,
				Message: swag.String(err.Error()),
			})
		}
		return operations.NewPostTaskCreated().WithPayload(result)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
