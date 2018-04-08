// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/organization-srv/proto/org/org.proto

/*
Package org is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/organization-srv/proto/org/org.proto

It has these top-level messages:
	Organization
	Member
	CreateRequest
	CreateResponse
	DeleteRequest
	DeleteResponse
	ReadRequest
	ReadResponse
	UpdateRequest
	UpdateResponse
	SearchRequest
	SearchResponse
	CreateMemberRequest
	CreateMemberResponse
	DeleteMemberRequest
	DeleteMemberResponse
	ReadMemberRequest
	ReadMemberResponse
	UpdateMemberRequest
	UpdateMemberResponse
	SearchMembersRequest
	SearchMembersResponse
*/
package org

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Org service

type OrgService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...client.CallOption) (*CreateMemberResponse, error)
	ReadMember(ctx context.Context, in *ReadMemberRequest, opts ...client.CallOption) (*ReadMemberResponse, error)
	DeleteMember(ctx context.Context, in *DeleteMemberRequest, opts ...client.CallOption) (*DeleteMemberResponse, error)
	UpdateMember(ctx context.Context, in *UpdateMemberRequest, opts ...client.CallOption) (*UpdateMemberResponse, error)
	SearchMembers(ctx context.Context, in *SearchMembersRequest, opts ...client.CallOption) (*SearchMembersResponse, error)
}

type orgService struct {
	c           client.Client
	serviceName string
}

func OrgServiceClient(serviceName string, c client.Client) OrgService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "org"
	}
	return &orgService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *orgService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Org.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Org.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Org.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Org.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Org.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...client.CallOption) (*CreateMemberResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Org.CreateMember", in)
	out := new(CreateMemberResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) ReadMember(ctx context.Context, in *ReadMemberRequest, opts ...client.CallOption) (*ReadMemberResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Org.ReadMember", in)
	out := new(ReadMemberResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) DeleteMember(ctx context.Context, in *DeleteMemberRequest, opts ...client.CallOption) (*DeleteMemberResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Org.DeleteMember", in)
	out := new(DeleteMemberResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) UpdateMember(ctx context.Context, in *UpdateMemberRequest, opts ...client.CallOption) (*UpdateMemberResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Org.UpdateMember", in)
	out := new(UpdateMemberResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgService) SearchMembers(ctx context.Context, in *SearchMembersRequest, opts ...client.CallOption) (*SearchMembersResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Org.SearchMembers", in)
	out := new(SearchMembersResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Org service

type OrgHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	CreateMember(context.Context, *CreateMemberRequest, *CreateMemberResponse) error
	ReadMember(context.Context, *ReadMemberRequest, *ReadMemberResponse) error
	DeleteMember(context.Context, *DeleteMemberRequest, *DeleteMemberResponse) error
	UpdateMember(context.Context, *UpdateMemberRequest, *UpdateMemberResponse) error
	SearchMembers(context.Context, *SearchMembersRequest, *SearchMembersResponse) error
}

func RegisterOrgHandler(s server.Server, hdlr OrgHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Org{hdlr}, opts...))
}

type Org struct {
	OrgHandler
}

func (h *Org) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.OrgHandler.Create(ctx, in, out)
}

func (h *Org) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.OrgHandler.Read(ctx, in, out)
}

func (h *Org) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.OrgHandler.Delete(ctx, in, out)
}

func (h *Org) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.OrgHandler.Update(ctx, in, out)
}

func (h *Org) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.OrgHandler.Search(ctx, in, out)
}

func (h *Org) CreateMember(ctx context.Context, in *CreateMemberRequest, out *CreateMemberResponse) error {
	return h.OrgHandler.CreateMember(ctx, in, out)
}

func (h *Org) ReadMember(ctx context.Context, in *ReadMemberRequest, out *ReadMemberResponse) error {
	return h.OrgHandler.ReadMember(ctx, in, out)
}

func (h *Org) DeleteMember(ctx context.Context, in *DeleteMemberRequest, out *DeleteMemberResponse) error {
	return h.OrgHandler.DeleteMember(ctx, in, out)
}

func (h *Org) UpdateMember(ctx context.Context, in *UpdateMemberRequest, out *UpdateMemberResponse) error {
	return h.OrgHandler.UpdateMember(ctx, in, out)
}

func (h *Org) SearchMembers(ctx context.Context, in *SearchMembersRequest, out *SearchMembersResponse) error {
	return h.OrgHandler.SearchMembers(ctx, in, out)
}
