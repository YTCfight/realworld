// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.19.1
// source: realworld/v1/realworld.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationGreeterAddComment = "/realworld.v1.Greeter/AddComment"
const OperationGreeterCreateArticle = "/realworld.v1.Greeter/CreateArticle"
const OperationGreeterDeleteArticle = "/realworld.v1.Greeter/DeleteArticle"
const OperationGreeterDeleteComment = "/realworld.v1.Greeter/DeleteComment"
const OperationGreeterFavoriteArticle = "/realworld.v1.Greeter/FavoriteArticle"
const OperationGreeterFeedArticles = "/realworld.v1.Greeter/FeedArticles"
const OperationGreeterFollowUser = "/realworld.v1.Greeter/FollowUser"
const OperationGreeterGetArticle = "/realworld.v1.Greeter/GetArticle"
const OperationGreeterGetComments = "/realworld.v1.Greeter/GetComments"
const OperationGreeterGetCurrentUser = "/realworld.v1.Greeter/GetCurrentUser"
const OperationGreeterGetProfile = "/realworld.v1.Greeter/GetProfile"
const OperationGreeterGetTags = "/realworld.v1.Greeter/GetTags"
const OperationGreeterListArticles = "/realworld.v1.Greeter/ListArticles"
const OperationGreeterLogin = "/realworld.v1.Greeter/Login"
const OperationGreeterRegister = "/realworld.v1.Greeter/Register"
const OperationGreeterUnFavoriteArticle = "/realworld.v1.Greeter/UnFavoriteArticle"
const OperationGreeterUnFollowUser = "/realworld.v1.Greeter/UnFollowUser"
const OperationGreeterUpdateArticle = "/realworld.v1.Greeter/UpdateArticle"
const OperationGreeterUpdateUser = "/realworld.v1.Greeter/UpdateUser"

type GreeterHTTPServer interface {
	AddComment(context.Context, *AddCommentRequest) (*SingleComment, error)
	CreateArticle(context.Context, *CreateArticleRequest) (*SingleArticleReply, error)
	DeleteArticle(context.Context, *DeleteArticleRequest) (*SingleArticleReply, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*SingleArticleReply, error)
	FavoriteArticle(context.Context, *FavoriteArticleRequest) (*ProfileReply, error)
	FeedArticles(context.Context, *FeedArticlesRequest) (*MultipleArticlesReply, error)
	FollowUser(context.Context, *FollowUserRequest) (*ProfileReply, error)
	GetArticle(context.Context, *GetArticleRequest) (*SingleArticleReply, error)
	GetComments(context.Context, *AddCommentRequest) (*MultipleCommentsReply, error)
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*UserReply, error)
	GetProfile(context.Context, *GetProfileRequest) (*ProfileReply, error)
	GetTags(context.Context, *GetTagsRequest) (*TagListReply, error)
	ListArticles(context.Context, *ListArticlesRequest) (*MultipleArticlesReply, error)
	Login(context.Context, *LoginRequest) (*UserReply, error)
	Register(context.Context, *RegisterRequest) (*UserReply, error)
	UnFavoriteArticle(context.Context, *UnFavoriteArticleRequest) (*ProfileReply, error)
	UnFollowUser(context.Context, *UnFollowUserRequest) (*ProfileReply, error)
	UpdateArticle(context.Context, *UpdateArticleRequest) (*SingleArticleReply, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UserReply, error)
}

func RegisterGreeterHTTPServer(s *http.Server, srv GreeterHTTPServer) {
	r := s.Route("/")
	r.POST("/api/users/login", _Greeter_Login0_HTTP_Handler(srv))
	r.POST("/api/users", _Greeter_Register0_HTTP_Handler(srv))
	r.GET("/api/user", _Greeter_GetCurrentUser0_HTTP_Handler(srv))
	r.PUT("/api/user", _Greeter_UpdateUser0_HTTP_Handler(srv))
	r.GET("/api/profiles/{username}", _Greeter_GetProfile0_HTTP_Handler(srv))
	r.POST("/api/profiles/{username}/follow", _Greeter_FollowUser0_HTTP_Handler(srv))
	r.DELETE("/api/profiles/{username}/follow", _Greeter_UnFollowUser0_HTTP_Handler(srv))
	r.GET("/api/articles", _Greeter_ListArticles0_HTTP_Handler(srv))
	r.GET("/api/articles/feed", _Greeter_FeedArticles0_HTTP_Handler(srv))
	r.GET("/api/articles/{slug}", _Greeter_GetArticle0_HTTP_Handler(srv))
	r.POST("/api/articles", _Greeter_CreateArticle0_HTTP_Handler(srv))
	r.PUT("/api/articles/{slug}", _Greeter_UpdateArticle0_HTTP_Handler(srv))
	r.DELETE("/api/articles/{slug}", _Greeter_DeleteArticle0_HTTP_Handler(srv))
	r.POST("/api/articles/{slug}/comments", _Greeter_AddComment0_HTTP_Handler(srv))
	r.POST("/api/articles/{slug}/comments", _Greeter_GetComments0_HTTP_Handler(srv))
	r.DELETE("/api/articles/{slug}/comments/{id}", _Greeter_DeleteComment0_HTTP_Handler(srv))
	r.POST("/api/articles/{slug}/favorite", _Greeter_FavoriteArticle0_HTTP_Handler(srv))
	r.DELETE("/api/articles/{slug}/favorite", _Greeter_UnFavoriteArticle0_HTTP_Handler(srv))
	r.GET("/api/tags", _Greeter_GetTags0_HTTP_Handler(srv))
}

func _Greeter_Login0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_Register0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_GetCurrentUser0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCurrentUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterGetCurrentUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_UpdateUser0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_GetProfile0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetProfileRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterGetProfile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetProfile(ctx, req.(*GetProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProfileReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_FollowUser0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FollowUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterFollowUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FollowUser(ctx, req.(*FollowUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProfileReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_UnFollowUser0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnFollowUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterUnFollowUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UnFollowUser(ctx, req.(*UnFollowUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProfileReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_ListArticles0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListArticlesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterListArticles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListArticles(ctx, req.(*ListArticlesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MultipleArticlesReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_FeedArticles0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FeedArticlesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterFeedArticles)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FeedArticles(ctx, req.(*FeedArticlesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MultipleArticlesReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_GetArticle0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterGetArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArticle(ctx, req.(*GetArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleArticleReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_CreateArticle0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterCreateArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateArticle(ctx, req.(*CreateArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleArticleReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_UpdateArticle0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterUpdateArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateArticle(ctx, req.(*UpdateArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleArticleReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_DeleteArticle0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterDeleteArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteArticle(ctx, req.(*DeleteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleArticleReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_AddComment0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterAddComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AddComment(ctx, req.(*AddCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleComment)
		return ctx.Result(200, reply)
	}
}

func _Greeter_GetComments0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddCommentRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterGetComments)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetComments(ctx, req.(*AddCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MultipleCommentsReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_DeleteComment0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteCommentRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterDeleteComment)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteComment(ctx, req.(*DeleteCommentRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SingleArticleReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_FavoriteArticle0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FavoriteArticleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterFavoriteArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FavoriteArticle(ctx, req.(*FavoriteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProfileReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_UnFavoriteArticle0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnFavoriteArticleRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterUnFavoriteArticle)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UnFavoriteArticle(ctx, req.(*UnFavoriteArticleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProfileReply)
		return ctx.Result(200, reply)
	}
}

func _Greeter_GetTags0_HTTP_Handler(srv GreeterHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetTagsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationGreeterGetTags)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetTags(ctx, req.(*GetTagsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TagListReply)
		return ctx.Result(200, reply)
	}
}

type GreeterHTTPClient interface {
	AddComment(ctx context.Context, req *AddCommentRequest, opts ...http.CallOption) (rsp *SingleComment, err error)
	CreateArticle(ctx context.Context, req *CreateArticleRequest, opts ...http.CallOption) (rsp *SingleArticleReply, err error)
	DeleteArticle(ctx context.Context, req *DeleteArticleRequest, opts ...http.CallOption) (rsp *SingleArticleReply, err error)
	DeleteComment(ctx context.Context, req *DeleteCommentRequest, opts ...http.CallOption) (rsp *SingleArticleReply, err error)
	FavoriteArticle(ctx context.Context, req *FavoriteArticleRequest, opts ...http.CallOption) (rsp *ProfileReply, err error)
	FeedArticles(ctx context.Context, req *FeedArticlesRequest, opts ...http.CallOption) (rsp *MultipleArticlesReply, err error)
	FollowUser(ctx context.Context, req *FollowUserRequest, opts ...http.CallOption) (rsp *ProfileReply, err error)
	GetArticle(ctx context.Context, req *GetArticleRequest, opts ...http.CallOption) (rsp *SingleArticleReply, err error)
	GetComments(ctx context.Context, req *AddCommentRequest, opts ...http.CallOption) (rsp *MultipleCommentsReply, err error)
	GetCurrentUser(ctx context.Context, req *GetCurrentUserRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	GetProfile(ctx context.Context, req *GetProfileRequest, opts ...http.CallOption) (rsp *ProfileReply, err error)
	GetTags(ctx context.Context, req *GetTagsRequest, opts ...http.CallOption) (rsp *TagListReply, err error)
	ListArticles(ctx context.Context, req *ListArticlesRequest, opts ...http.CallOption) (rsp *MultipleArticlesReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *UserReply, err error)
	UnFavoriteArticle(ctx context.Context, req *UnFavoriteArticleRequest, opts ...http.CallOption) (rsp *ProfileReply, err error)
	UnFollowUser(ctx context.Context, req *UnFollowUserRequest, opts ...http.CallOption) (rsp *ProfileReply, err error)
	UpdateArticle(ctx context.Context, req *UpdateArticleRequest, opts ...http.CallOption) (rsp *SingleArticleReply, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UserReply, err error)
}

type GreeterHTTPClientImpl struct {
	cc *http.Client
}

func NewGreeterHTTPClient(client *http.Client) GreeterHTTPClient {
	return &GreeterHTTPClientImpl{client}
}

func (c *GreeterHTTPClientImpl) AddComment(ctx context.Context, in *AddCommentRequest, opts ...http.CallOption) (*SingleComment, error) {
	var out SingleComment
	pattern := "/api/articles/{slug}/comments"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGreeterAddComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...http.CallOption) (*SingleArticleReply, error) {
	var out SingleArticleReply
	pattern := "/api/articles"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGreeterCreateArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...http.CallOption) (*SingleArticleReply, error) {
	var out SingleArticleReply
	pattern := "/api/articles/{slug}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterDeleteArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...http.CallOption) (*SingleArticleReply, error) {
	var out SingleArticleReply
	pattern := "/api/articles/{slug}/comments/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterDeleteComment))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) FavoriteArticle(ctx context.Context, in *FavoriteArticleRequest, opts ...http.CallOption) (*ProfileReply, error) {
	var out ProfileReply
	pattern := "/api/articles/{slug}/favorite"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGreeterFavoriteArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) FeedArticles(ctx context.Context, in *FeedArticlesRequest, opts ...http.CallOption) (*MultipleArticlesReply, error) {
	var out MultipleArticlesReply
	pattern := "/api/articles/feed"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterFeedArticles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) FollowUser(ctx context.Context, in *FollowUserRequest, opts ...http.CallOption) (*ProfileReply, error) {
	var out ProfileReply
	pattern := "/api/profiles/{username}/follow"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGreeterFollowUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...http.CallOption) (*SingleArticleReply, error) {
	var out SingleArticleReply
	pattern := "/api/articles/{slug}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterGetArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) GetComments(ctx context.Context, in *AddCommentRequest, opts ...http.CallOption) (*MultipleCommentsReply, error) {
	var out MultipleCommentsReply
	pattern := "/api/articles/{slug}/comments"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGreeterGetComments))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/api/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterGetCurrentUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) GetProfile(ctx context.Context, in *GetProfileRequest, opts ...http.CallOption) (*ProfileReply, error) {
	var out ProfileReply
	pattern := "/api/profiles/{username}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterGetProfile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) GetTags(ctx context.Context, in *GetTagsRequest, opts ...http.CallOption) (*TagListReply, error) {
	var out TagListReply
	pattern := "/api/tags"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterGetTags))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) ListArticles(ctx context.Context, in *ListArticlesRequest, opts ...http.CallOption) (*MultipleArticlesReply, error) {
	var out MultipleArticlesReply
	pattern := "/api/articles"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterListArticles))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/api/users/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGreeterLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/api/users"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGreeterRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) UnFavoriteArticle(ctx context.Context, in *UnFavoriteArticleRequest, opts ...http.CallOption) (*ProfileReply, error) {
	var out ProfileReply
	pattern := "/api/articles/{slug}/favorite"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterUnFavoriteArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) UnFollowUser(ctx context.Context, in *UnFollowUserRequest, opts ...http.CallOption) (*ProfileReply, error) {
	var out ProfileReply
	pattern := "/api/profiles/{username}/follow"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationGreeterUnFollowUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...http.CallOption) (*SingleArticleReply, error) {
	var out SingleArticleReply
	pattern := "/api/articles/{slug}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGreeterUpdateArticle))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GreeterHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UserReply, error) {
	var out UserReply
	pattern := "/api/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationGreeterUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}