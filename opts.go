package anthropic

import (
	"github.com/3JoB/resty-ilo"
	"github.com/3JoB/ulib/err"
	// "github.com/google/uuid"

	"github.com/3JoB/anthropic-sdk-go/v2/context"
	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
)

type Opts struct {
	Message   data.MessageModule // Chunked message structure
	ContextID string             // Session ID. If empty, a new session is automatically created. If not empty, an attempt is made to find an existing session.
	Sender    resp.Sender
}

func (opts *Opts) newCtx() *context.Context {
	return &context.Context{
		Response: &resp.Response{},
		Human:    opts.Message.Human,
	}
}

// Send data to the API endpoint. Before sending out, the data will be processed into a form that the API can recognize.
//
// This method will be used to handle context requests.
//
// The context parameter comes from *Context.CtxData, please do not modify or process it by yourself, the context will be automatically processed when the previous request is executed.
/*func (ah *AnthropicClient) SendWithContext(sender *Sender, context string) (ctx *Context, err error) {
	if err := ah.c(sender); err != nil {
		return nil, err
	}
	sender.Prompt, err = addPrompt(context, sender.Prompt)
	if err != nil {
		return nil, err
	}
	return sender.Complete(ah.client)
}*/

// Make a processed request to an API endpoint.
func (req *Opts) Complete(ctx *context.Context, client *resty.Client) (*context.Context, error) {
	rq := client.R().SetBody(req.Sender)
	r, errs := rq.Post(APIComplete)
	if errs != nil {
		return ctx, &err.Err{Op: "request", Err: errs.Error()}
	}
	defer r.RawBody().Close()

	ctx.ID = req.ContextID
	if errs := r.Bind(ctx.Response); errs != nil {
		return ctx, &err.Err{Op: "request", E: errs}
	}

	ctx.RawData = r.String()

	if !r.IsStatusCode(200) {
		return ctx, &err.Err{Op: "request", Err: ctx.RawData}
	}

	req.Message.Assistant = ctx.Response.Completion

	if !ctx.Add() {
		return ctx, &err.Err{Op: "request", Err: "Add failed"}
	}

	return ctx, nil
}
