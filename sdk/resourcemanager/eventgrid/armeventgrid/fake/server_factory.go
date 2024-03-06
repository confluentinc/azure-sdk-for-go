//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armeventgrid.ClientFactory type.
type ServerFactory struct {
	CaCertificatesServer                         CaCertificatesServer
	ChannelsServer                               ChannelsServer
	ClientGroupsServer                           ClientGroupsServer
	ClientsServer                                ClientsServer
	DomainEventSubscriptionsServer               DomainEventSubscriptionsServer
	DomainTopicEventSubscriptionsServer          DomainTopicEventSubscriptionsServer
	DomainTopicsServer                           DomainTopicsServer
	DomainsServer                                DomainsServer
	EventSubscriptionsServer                     EventSubscriptionsServer
	ExtensionTopicsServer                        ExtensionTopicsServer
	NamespaceTopicEventSubscriptionsServer       NamespaceTopicEventSubscriptionsServer
	NamespaceTopicsServer                        NamespaceTopicsServer
	NamespacesServer                             NamespacesServer
	NetworkSecurityPerimeterConfigurationsServer NetworkSecurityPerimeterConfigurationsServer
	OperationsServer                             OperationsServer
	PartnerConfigurationsServer                  PartnerConfigurationsServer
	PartnerDestinationsServer                    PartnerDestinationsServer
	PartnerNamespacesServer                      PartnerNamespacesServer
	PartnerRegistrationsServer                   PartnerRegistrationsServer
	PartnerTopicEventSubscriptionsServer         PartnerTopicEventSubscriptionsServer
	PartnerTopicsServer                          PartnerTopicsServer
	PermissionBindingsServer                     PermissionBindingsServer
	PrivateEndpointConnectionsServer             PrivateEndpointConnectionsServer
	PrivateLinkResourcesServer                   PrivateLinkResourcesServer
	SystemTopicEventSubscriptionsServer          SystemTopicEventSubscriptionsServer
	SystemTopicsServer                           SystemTopicsServer
	TopicEventSubscriptionsServer                TopicEventSubscriptionsServer
	TopicSpacesServer                            TopicSpacesServer
	TopicTypesServer                             TopicTypesServer
	TopicsServer                                 TopicsServer
	VerifiedPartnersServer                       VerifiedPartnersServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armeventgrid.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armeventgrid.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                            *ServerFactory
	trMu                                           sync.Mutex
	trCaCertificatesServer                         *CaCertificatesServerTransport
	trChannelsServer                               *ChannelsServerTransport
	trClientGroupsServer                           *ClientGroupsServerTransport
	trClientsServer                                *ClientsServerTransport
	trDomainEventSubscriptionsServer               *DomainEventSubscriptionsServerTransport
	trDomainTopicEventSubscriptionsServer          *DomainTopicEventSubscriptionsServerTransport
	trDomainTopicsServer                           *DomainTopicsServerTransport
	trDomainsServer                                *DomainsServerTransport
	trEventSubscriptionsServer                     *EventSubscriptionsServerTransport
	trExtensionTopicsServer                        *ExtensionTopicsServerTransport
	trNamespaceTopicEventSubscriptionsServer       *NamespaceTopicEventSubscriptionsServerTransport
	trNamespaceTopicsServer                        *NamespaceTopicsServerTransport
	trNamespacesServer                             *NamespacesServerTransport
	trNetworkSecurityPerimeterConfigurationsServer *NetworkSecurityPerimeterConfigurationsServerTransport
	trOperationsServer                             *OperationsServerTransport
	trPartnerConfigurationsServer                  *PartnerConfigurationsServerTransport
	trPartnerDestinationsServer                    *PartnerDestinationsServerTransport
	trPartnerNamespacesServer                      *PartnerNamespacesServerTransport
	trPartnerRegistrationsServer                   *PartnerRegistrationsServerTransport
	trPartnerTopicEventSubscriptionsServer         *PartnerTopicEventSubscriptionsServerTransport
	trPartnerTopicsServer                          *PartnerTopicsServerTransport
	trPermissionBindingsServer                     *PermissionBindingsServerTransport
	trPrivateEndpointConnectionsServer             *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer                   *PrivateLinkResourcesServerTransport
	trSystemTopicEventSubscriptionsServer          *SystemTopicEventSubscriptionsServerTransport
	trSystemTopicsServer                           *SystemTopicsServerTransport
	trTopicEventSubscriptionsServer                *TopicEventSubscriptionsServerTransport
	trTopicSpacesServer                            *TopicSpacesServerTransport
	trTopicTypesServer                             *TopicTypesServerTransport
	trTopicsServer                                 *TopicsServerTransport
	trVerifiedPartnersServer                       *VerifiedPartnersServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "CaCertificatesClient":
		initServer(s, &s.trCaCertificatesServer, func() *CaCertificatesServerTransport {
			return NewCaCertificatesServerTransport(&s.srv.CaCertificatesServer)
		})
		resp, err = s.trCaCertificatesServer.Do(req)
	case "ChannelsClient":
		initServer(s, &s.trChannelsServer, func() *ChannelsServerTransport { return NewChannelsServerTransport(&s.srv.ChannelsServer) })
		resp, err = s.trChannelsServer.Do(req)
	case "ClientGroupsClient":
		initServer(s, &s.trClientGroupsServer, func() *ClientGroupsServerTransport { return NewClientGroupsServerTransport(&s.srv.ClientGroupsServer) })
		resp, err = s.trClientGroupsServer.Do(req)
	case "ClientsClient":
		initServer(s, &s.trClientsServer, func() *ClientsServerTransport { return NewClientsServerTransport(&s.srv.ClientsServer) })
		resp, err = s.trClientsServer.Do(req)
	case "DomainEventSubscriptionsClient":
		initServer(s, &s.trDomainEventSubscriptionsServer, func() *DomainEventSubscriptionsServerTransport {
			return NewDomainEventSubscriptionsServerTransport(&s.srv.DomainEventSubscriptionsServer)
		})
		resp, err = s.trDomainEventSubscriptionsServer.Do(req)
	case "DomainTopicEventSubscriptionsClient":
		initServer(s, &s.trDomainTopicEventSubscriptionsServer, func() *DomainTopicEventSubscriptionsServerTransport {
			return NewDomainTopicEventSubscriptionsServerTransport(&s.srv.DomainTopicEventSubscriptionsServer)
		})
		resp, err = s.trDomainTopicEventSubscriptionsServer.Do(req)
	case "DomainTopicsClient":
		initServer(s, &s.trDomainTopicsServer, func() *DomainTopicsServerTransport { return NewDomainTopicsServerTransport(&s.srv.DomainTopicsServer) })
		resp, err = s.trDomainTopicsServer.Do(req)
	case "DomainsClient":
		initServer(s, &s.trDomainsServer, func() *DomainsServerTransport { return NewDomainsServerTransport(&s.srv.DomainsServer) })
		resp, err = s.trDomainsServer.Do(req)
	case "EventSubscriptionsClient":
		initServer(s, &s.trEventSubscriptionsServer, func() *EventSubscriptionsServerTransport {
			return NewEventSubscriptionsServerTransport(&s.srv.EventSubscriptionsServer)
		})
		resp, err = s.trEventSubscriptionsServer.Do(req)
	case "ExtensionTopicsClient":
		initServer(s, &s.trExtensionTopicsServer, func() *ExtensionTopicsServerTransport {
			return NewExtensionTopicsServerTransport(&s.srv.ExtensionTopicsServer)
		})
		resp, err = s.trExtensionTopicsServer.Do(req)
	case "NamespaceTopicEventSubscriptionsClient":
		initServer(s, &s.trNamespaceTopicEventSubscriptionsServer, func() *NamespaceTopicEventSubscriptionsServerTransport {
			return NewNamespaceTopicEventSubscriptionsServerTransport(&s.srv.NamespaceTopicEventSubscriptionsServer)
		})
		resp, err = s.trNamespaceTopicEventSubscriptionsServer.Do(req)
	case "NamespaceTopicsClient":
		initServer(s, &s.trNamespaceTopicsServer, func() *NamespaceTopicsServerTransport {
			return NewNamespaceTopicsServerTransport(&s.srv.NamespaceTopicsServer)
		})
		resp, err = s.trNamespaceTopicsServer.Do(req)
	case "NamespacesClient":
		initServer(s, &s.trNamespacesServer, func() *NamespacesServerTransport { return NewNamespacesServerTransport(&s.srv.NamespacesServer) })
		resp, err = s.trNamespacesServer.Do(req)
	case "NetworkSecurityPerimeterConfigurationsClient":
		initServer(s, &s.trNetworkSecurityPerimeterConfigurationsServer, func() *NetworkSecurityPerimeterConfigurationsServerTransport {
			return NewNetworkSecurityPerimeterConfigurationsServerTransport(&s.srv.NetworkSecurityPerimeterConfigurationsServer)
		})
		resp, err = s.trNetworkSecurityPerimeterConfigurationsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PartnerConfigurationsClient":
		initServer(s, &s.trPartnerConfigurationsServer, func() *PartnerConfigurationsServerTransport {
			return NewPartnerConfigurationsServerTransport(&s.srv.PartnerConfigurationsServer)
		})
		resp, err = s.trPartnerConfigurationsServer.Do(req)
	case "PartnerDestinationsClient":
		initServer(s, &s.trPartnerDestinationsServer, func() *PartnerDestinationsServerTransport {
			return NewPartnerDestinationsServerTransport(&s.srv.PartnerDestinationsServer)
		})
		resp, err = s.trPartnerDestinationsServer.Do(req)
	case "PartnerNamespacesClient":
		initServer(s, &s.trPartnerNamespacesServer, func() *PartnerNamespacesServerTransport {
			return NewPartnerNamespacesServerTransport(&s.srv.PartnerNamespacesServer)
		})
		resp, err = s.trPartnerNamespacesServer.Do(req)
	case "PartnerRegistrationsClient":
		initServer(s, &s.trPartnerRegistrationsServer, func() *PartnerRegistrationsServerTransport {
			return NewPartnerRegistrationsServerTransport(&s.srv.PartnerRegistrationsServer)
		})
		resp, err = s.trPartnerRegistrationsServer.Do(req)
	case "PartnerTopicEventSubscriptionsClient":
		initServer(s, &s.trPartnerTopicEventSubscriptionsServer, func() *PartnerTopicEventSubscriptionsServerTransport {
			return NewPartnerTopicEventSubscriptionsServerTransport(&s.srv.PartnerTopicEventSubscriptionsServer)
		})
		resp, err = s.trPartnerTopicEventSubscriptionsServer.Do(req)
	case "PartnerTopicsClient":
		initServer(s, &s.trPartnerTopicsServer, func() *PartnerTopicsServerTransport {
			return NewPartnerTopicsServerTransport(&s.srv.PartnerTopicsServer)
		})
		resp, err = s.trPartnerTopicsServer.Do(req)
	case "PermissionBindingsClient":
		initServer(s, &s.trPermissionBindingsServer, func() *PermissionBindingsServerTransport {
			return NewPermissionBindingsServerTransport(&s.srv.PermissionBindingsServer)
		})
		resp, err = s.trPermissionBindingsServer.Do(req)
	case "PrivateEndpointConnectionsClient":
		initServer(s, &s.trPrivateEndpointConnectionsServer, func() *PrivateEndpointConnectionsServerTransport {
			return NewPrivateEndpointConnectionsServerTransport(&s.srv.PrivateEndpointConnectionsServer)
		})
		resp, err = s.trPrivateEndpointConnectionsServer.Do(req)
	case "PrivateLinkResourcesClient":
		initServer(s, &s.trPrivateLinkResourcesServer, func() *PrivateLinkResourcesServerTransport {
			return NewPrivateLinkResourcesServerTransport(&s.srv.PrivateLinkResourcesServer)
		})
		resp, err = s.trPrivateLinkResourcesServer.Do(req)
	case "SystemTopicEventSubscriptionsClient":
		initServer(s, &s.trSystemTopicEventSubscriptionsServer, func() *SystemTopicEventSubscriptionsServerTransport {
			return NewSystemTopicEventSubscriptionsServerTransport(&s.srv.SystemTopicEventSubscriptionsServer)
		})
		resp, err = s.trSystemTopicEventSubscriptionsServer.Do(req)
	case "SystemTopicsClient":
		initServer(s, &s.trSystemTopicsServer, func() *SystemTopicsServerTransport { return NewSystemTopicsServerTransport(&s.srv.SystemTopicsServer) })
		resp, err = s.trSystemTopicsServer.Do(req)
	case "TopicEventSubscriptionsClient":
		initServer(s, &s.trTopicEventSubscriptionsServer, func() *TopicEventSubscriptionsServerTransport {
			return NewTopicEventSubscriptionsServerTransport(&s.srv.TopicEventSubscriptionsServer)
		})
		resp, err = s.trTopicEventSubscriptionsServer.Do(req)
	case "TopicSpacesClient":
		initServer(s, &s.trTopicSpacesServer, func() *TopicSpacesServerTransport { return NewTopicSpacesServerTransport(&s.srv.TopicSpacesServer) })
		resp, err = s.trTopicSpacesServer.Do(req)
	case "TopicTypesClient":
		initServer(s, &s.trTopicTypesServer, func() *TopicTypesServerTransport { return NewTopicTypesServerTransport(&s.srv.TopicTypesServer) })
		resp, err = s.trTopicTypesServer.Do(req)
	case "TopicsClient":
		initServer(s, &s.trTopicsServer, func() *TopicsServerTransport { return NewTopicsServerTransport(&s.srv.TopicsServer) })
		resp, err = s.trTopicsServer.Do(req)
	case "VerifiedPartnersClient":
		initServer(s, &s.trVerifiedPartnersServer, func() *VerifiedPartnersServerTransport {
			return NewVerifiedPartnersServerTransport(&s.srv.VerifiedPartnersServer)
		})
		resp, err = s.trVerifiedPartnersServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
