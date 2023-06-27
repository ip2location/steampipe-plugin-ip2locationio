package ip2locationio

import (
	"context"
	"errors"
	ip2location "github.com/ip2location/ip2location-io-go/ip2locationio"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"os"
)

func connectGeolocation(_ context.Context, d *plugin.QueryData) (*ip2location.IPGeolocation, error) {

	// Load connection from cache
	cacheKey := "ip2locationiogeolocation"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*ip2location.IPGeolocation), nil
	}

	// Default to the env var setting
	apiKey := os.Getenv("IP2LOCATIONIO_API_KEY")

	// Prefer config settings
	ip2locationioConfig := GetConfig(d.Connection)
	if ip2locationioConfig.ApiKey != nil {
		apiKey = *ip2locationioConfig.ApiKey
	}

	// Error if the minimum config is not set
	if apiKey == "" {
		return nil, errors.New("API key must be configured")
	}

	config, err := ip2location.OpenConfiguration(apiKey)

	if err != nil {
		return nil, err
	}
	ipl, err := ip2location.OpenIPGeolocation(config)

	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, ipl)

	return ipl, nil
}

func connectWhois(_ context.Context, d *plugin.QueryData) (*ip2location.DomainWhois, error) {

	// Load connection from cache
	cacheKey := "ip2locationiowhois"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*ip2location.DomainWhois), nil
	}

	// Default to the env var setting
	apiKey := os.Getenv("IP2LOCATIONIO_API_KEY")

	// Prefer config settings
	ip2locationioConfig := GetConfig(d.Connection)
	if ip2locationioConfig.ApiKey != nil {
		apiKey = *ip2locationioConfig.ApiKey
	}

	// Error if the minimum config is not set
	if apiKey == "" {
		return nil, errors.New("API key must be configured")
	}

	config, err := ip2location.OpenConfiguration(apiKey)

	if err != nil {
		return nil, err
	}
	whois, err := ip2location.OpenDomainWhois(config)

	if err != nil {
		return nil, err
	}

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, whois)

	return whois, nil
}
