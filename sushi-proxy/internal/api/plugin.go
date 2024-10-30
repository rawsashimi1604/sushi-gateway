package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/constant"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/db"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/gateway"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/model"
	"github.com/rawsashimi1604/sushi-gateway/sushi-proxy/internal/util"
	"log/slog"
	"net/http"
	"strings"
)

type PluginController struct {
	pluginRepo *db.PluginRepository
}

func NewPluginController(pluginRepo *db.PluginRepository) *PluginController {
	return &PluginController{pluginRepo: pluginRepo}
}

func (p *PluginController) RegisterRoutes(router *mux.Router) {
	router.Path("/").Methods("PUT").Handler(
		ProtectRouteUsingJWT(
			ProtectRouteWhenUsingDblessMode(p.UpdatePlugin())))
	router.Path("/").Methods("POST").Handler(
		ProtectRouteUsingJWT(
			ProtectRouteWhenUsingDblessMode(p.AddPlugin())))
	router.Path("/").Methods("DELETE").Handler(
		ProtectRouteUsingJWT(
			ProtectRouteWhenUsingDblessMode(p.DeletePlugin())))
}

// PluginDTO represents the structure for adding a new route, including the service name.
type PluginDTO struct {
	Scope  string             `json:"scope"`
	Name   string             `json:"name"`
	Plugin model.PluginConfig `json:"plugin"`
}

/*
	As a developer i want to update plugin.

	update plugin to global configuration
	update plugin to service
	update plugin to route

	add new plugin to global configuration
	add new plugin to service
	add new plugin to route

	delete plugin from global configuration
	delete plugin from service
	delete plugin from route.
*/

// TODO: create these routes.
func (p *PluginController) UpdatePlugin() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		// Decode the request
		var pluginDTO PluginDTO
		if err := json.NewDecoder(req.Body).Decode(&pluginDTO); err != nil {
			slog.Info("Failed to decode plugin DTO from request: " + err.Error())
			httperr := &model.HttpError{
				Code:     "UPDATE_PLUGIN_ERR",
				Message:  "failed to decode plugin from request body",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Validate the DTO contents
		if strings.ToLower(pluginDTO.Scope) != "global" &&
			strings.ToLower(pluginDTO.Scope) != "service" &&
			strings.ToLower(pluginDTO.Scope) != "route" {
			httperr := &model.HttpError{
				Code:     "UPDATE_PLUGIN_ERR",
				Message:  "scope is required and must be global, service or route",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// If the scope is service or route,
		// then we must check for a service name or route name to update the plugin to.
		if pluginDTO.Scope != "global" && pluginDTO.Name == "" {
			httperr := &model.HttpError{
				Code:     "UPDATE_PLUGIN_ERR",
				Message:  "name is required when we are adding plugin to service or route scope.",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Validate that the plugin exist already
		checkPlugins, err := p.pluginRepo.GetPlugins(pluginDTO.Scope, pluginDTO.Name)
		if err != nil {
			slog.Info("failed to get plugins")
			slog.Info(err.Error())
			httperr := &model.HttpError{
				Code:     "UPDATE_PLUGIN_ERR",
				Message:  "failed to add plugin into database",
				HttpCode: http.StatusInternalServerError,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		found := false
		var foundPlugin model.PluginConfig
		for _, pluginToCheck := range checkPlugins {
			if pluginToCheck.Name == pluginDTO.Plugin.Name {
				foundPlugin = pluginToCheck
				found = true
				break
			}
		}
		if !found {
			httperr := &model.HttpError{
				Code:     "UPDATE_PLUGIN_ERR",
				Message:  "plugin not found",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Update the plugin model
		foundPlugin.Config = pluginDTO.Plugin.Config
		foundPlugin.Enabled = pluginDTO.Plugin.Enabled

		err = p.pluginRepo.UpdatePlugin(pluginDTO.Scope, foundPlugin)
		if err != nil {
			slog.Info("failed to update plugin")
			slog.Info(err.Error())
			httperr := &model.HttpError{
				Code:     "UPDATE_PLUGIN_ERR",
				Message:  "failed to add plugin into database",
				HttpCode: http.StatusInternalServerError,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Send a success response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Plugin updated successfully",
		})
	}
}

func (p *PluginController) AddPlugin() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		// Decode the request
		var pluginDTO PluginDTO
		if err := json.NewDecoder(req.Body).Decode(&pluginDTO); err != nil {
			slog.Info("Failed to decode plugin DTO from request: " + err.Error())
			httperr := &model.HttpError{
				Code:     "CREATE_PLUGIN_ERR",
				Message:  "failed to decode plugin from request body",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Validate the DTO contents
		if strings.ToLower(pluginDTO.Scope) != "global" &&
			strings.ToLower(pluginDTO.Scope) != "service" &&
			strings.ToLower(pluginDTO.Scope) != "route" {
			httperr := &model.HttpError{
				Code:     "CREATE_PLUGIN_ERR",
				Message:  "scope is required and must be global, service or route",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// If the scope is service or route,
		// then we must check for a service name or route name to add the plugin to.
		if pluginDTO.Scope != "global" && pluginDTO.Name == "" {
			httperr := &model.HttpError{
				Code:     "CREATE_PLUGIN_ERR",
				Message:  "name is required when we are adding plugin to service or route scope.",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Validate that plugin is one of the available plugins
		if !util.SliceContainsString(constant.AVAILABLE_PLUGINS, pluginDTO.Plugin.Name) {
			httperr := &model.HttpError{
				Code:     "CREATE_PLUGIN_ERR",
				Message:  "plugin configuration name is not available: " + pluginDTO.Plugin.Name,
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Validate that the plugin does not exist already
		checkPlugins, err := p.pluginRepo.GetPlugins(pluginDTO.Scope, pluginDTO.Name)
		if err != nil {
			slog.Info("failed to get plugins")
			slog.Info(err.Error())
			httperr := &model.HttpError{
				Code:     "CREATE_PLUGIN_ERR",
				Message:  "failed to add plugin into database",
				HttpCode: http.StatusInternalServerError,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		for _, pluginToCheck := range checkPlugins {
			if pluginToCheck.Name == pluginDTO.Plugin.Name {
				httperr := &model.HttpError{
					Code:     "CREATE_PLUGIN_ERR",
					Message:  "plugin already exists.",
					HttpCode: http.StatusInternalServerError,
				}
				httperr.WriteLogMessage()
				httperr.WriteJSONResponse(w)
				return
			}
		}

		// Inject UUID into plugin
		uuidGenerator := gateway.NewUUIDGenerator()
		uuidGenerator.GenerateUUIDForPlugin(&pluginDTO.Plugin)

		// Add the plugin
		err = p.pluginRepo.AddPlugin(pluginDTO.Scope, pluginDTO.Plugin, pluginDTO.Name)
		if err != nil {
			slog.Info(err.Error())
			httperr := &model.HttpError{
				Code:     "CREATE_PLUGIN_ERR",
				Message:  "failed to add plugin into database",
				HttpCode: http.StatusInternalServerError,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Send a success response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Plugin created successfully",
		})
	}
}

func (p *PluginController) DeletePlugin() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		// Decode the request body to get the scope and plugin details
		var pluginDTO PluginDTO
		if err := json.NewDecoder(req.Body).Decode(&pluginDTO); err != nil {
			slog.Info("Failed to decode plugin DTO from request: " + err.Error())
			httperr := &model.HttpError{
				Code:     "DELETE_PLUGIN_ERR",
				Message:  "failed to decode plugin from request body",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Validate the scope
		if strings.ToLower(pluginDTO.Scope) != "global" &&
			strings.ToLower(pluginDTO.Scope) != "service" &&
			strings.ToLower(pluginDTO.Scope) != "route" {
			httperr := &model.HttpError{
				Code:     "DELETE_PLUGIN_ERR",
				Message:  "scope is required and must be global, service, or route",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Check that the plugin name is provided if it’s service or route scoped
		if pluginDTO.Scope != "global" && pluginDTO.Name == "" {
			httperr := &model.HttpError{
				Code:     "DELETE_PLUGIN_ERR",
				Message:  "name is required when deleting plugin at service or route scope",
				HttpCode: http.StatusBadRequest,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Call the repository to delete the plugin based on scope
		err := p.pluginRepo.DeletePlugin(pluginDTO.Scope, pluginDTO.Plugin.Name, pluginDTO.Name)
		if err != nil {
			slog.Info("failed to delete plugin: " + err.Error())
			httperr := &model.HttpError{
				Code:     "DELETE_PLUGIN_ERR",
				Message:  "failed to delete plugin from the database",
				HttpCode: http.StatusInternalServerError,
			}
			httperr.WriteLogMessage()
			httperr.WriteJSONResponse(w)
			return
		}

		// Send a success response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Plugin deleted successfully",
		})
	}
}
