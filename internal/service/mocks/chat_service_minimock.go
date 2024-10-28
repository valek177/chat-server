// Code generated by http://github.com/gojuno/minimock (v3.4.1). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/valek177/chat-server/internal/service.ChatService -o chat_service_minimock.go -n ChatServiceMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/valek177/chat-server/grpc/pkg/chat_v1"
)

// ChatServiceMock implements mm_service.ChatService
type ChatServiceMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreateChat          func(ctx context.Context, req *chat_v1.CreateChatRequest) (i1 int64, err error)
	funcCreateChatOrigin    string
	inspectFuncCreateChat   func(ctx context.Context, req *chat_v1.CreateChatRequest)
	afterCreateChatCounter  uint64
	beforeCreateChatCounter uint64
	CreateChatMock          mChatServiceMockCreateChat

	funcDeleteChat          func(ctx context.Context, id int64) (err error)
	funcDeleteChatOrigin    string
	inspectFuncDeleteChat   func(ctx context.Context, id int64)
	afterDeleteChatCounter  uint64
	beforeDeleteChatCounter uint64
	DeleteChatMock          mChatServiceMockDeleteChat
}

// NewChatServiceMock returns a mock for mm_service.ChatService
func NewChatServiceMock(t minimock.Tester) *ChatServiceMock {
	m := &ChatServiceMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateChatMock = mChatServiceMockCreateChat{mock: m}
	m.CreateChatMock.callArgs = []*ChatServiceMockCreateChatParams{}

	m.DeleteChatMock = mChatServiceMockDeleteChat{mock: m}
	m.DeleteChatMock.callArgs = []*ChatServiceMockDeleteChatParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mChatServiceMockCreateChat struct {
	optional           bool
	mock               *ChatServiceMock
	defaultExpectation *ChatServiceMockCreateChatExpectation
	expectations       []*ChatServiceMockCreateChatExpectation

	callArgs []*ChatServiceMockCreateChatParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// ChatServiceMockCreateChatExpectation specifies expectation struct of the ChatService.CreateChat
type ChatServiceMockCreateChatExpectation struct {
	mock               *ChatServiceMock
	params             *ChatServiceMockCreateChatParams
	paramPtrs          *ChatServiceMockCreateChatParamPtrs
	expectationOrigins ChatServiceMockCreateChatExpectationOrigins
	results            *ChatServiceMockCreateChatResults
	returnOrigin       string
	Counter            uint64
}

// ChatServiceMockCreateChatParams contains parameters of the ChatService.CreateChat
type ChatServiceMockCreateChatParams struct {
	ctx context.Context
	req *chat_v1.CreateChatRequest
}

// ChatServiceMockCreateChatParamPtrs contains pointers to parameters of the ChatService.CreateChat
type ChatServiceMockCreateChatParamPtrs struct {
	ctx *context.Context
	req **chat_v1.CreateChatRequest
}

// ChatServiceMockCreateChatResults contains results of the ChatService.CreateChat
type ChatServiceMockCreateChatResults struct {
	i1  int64
	err error
}

// ChatServiceMockCreateChatOrigins contains origins of expectations of the ChatService.CreateChat
type ChatServiceMockCreateChatExpectationOrigins struct {
	origin    string
	originCtx string
	originReq string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCreateChat *mChatServiceMockCreateChat) Optional() *mChatServiceMockCreateChat {
	mmCreateChat.optional = true
	return mmCreateChat
}

// Expect sets up expected params for ChatService.CreateChat
func (mmCreateChat *mChatServiceMockCreateChat) Expect(ctx context.Context, req *chat_v1.CreateChatRequest) *mChatServiceMockCreateChat {
	if mmCreateChat.mock.funcCreateChat != nil {
		mmCreateChat.mock.t.Fatalf("ChatServiceMock.CreateChat mock is already set by Set")
	}

	if mmCreateChat.defaultExpectation == nil {
		mmCreateChat.defaultExpectation = &ChatServiceMockCreateChatExpectation{}
	}

	if mmCreateChat.defaultExpectation.paramPtrs != nil {
		mmCreateChat.mock.t.Fatalf("ChatServiceMock.CreateChat mock is already set by ExpectParams functions")
	}

	mmCreateChat.defaultExpectation.params = &ChatServiceMockCreateChatParams{ctx, req}
	mmCreateChat.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmCreateChat.expectations {
		if minimock.Equal(e.params, mmCreateChat.defaultExpectation.params) {
			mmCreateChat.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateChat.defaultExpectation.params)
		}
	}

	return mmCreateChat
}

// ExpectCtxParam1 sets up expected param ctx for ChatService.CreateChat
func (mmCreateChat *mChatServiceMockCreateChat) ExpectCtxParam1(ctx context.Context) *mChatServiceMockCreateChat {
	if mmCreateChat.mock.funcCreateChat != nil {
		mmCreateChat.mock.t.Fatalf("ChatServiceMock.CreateChat mock is already set by Set")
	}

	if mmCreateChat.defaultExpectation == nil {
		mmCreateChat.defaultExpectation = &ChatServiceMockCreateChatExpectation{}
	}

	if mmCreateChat.defaultExpectation.params != nil {
		mmCreateChat.mock.t.Fatalf("ChatServiceMock.CreateChat mock is already set by Expect")
	}

	if mmCreateChat.defaultExpectation.paramPtrs == nil {
		mmCreateChat.defaultExpectation.paramPtrs = &ChatServiceMockCreateChatParamPtrs{}
	}
	mmCreateChat.defaultExpectation.paramPtrs.ctx = &ctx
	mmCreateChat.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmCreateChat
}

// ExpectReqParam2 sets up expected param req for ChatService.CreateChat
func (mmCreateChat *mChatServiceMockCreateChat) ExpectReqParam2(req *chat_v1.CreateChatRequest) *mChatServiceMockCreateChat {
	if mmCreateChat.mock.funcCreateChat != nil {
		mmCreateChat.mock.t.Fatalf("ChatServiceMock.CreateChat mock is already set by Set")
	}

	if mmCreateChat.defaultExpectation == nil {
		mmCreateChat.defaultExpectation = &ChatServiceMockCreateChatExpectation{}
	}

	if mmCreateChat.defaultExpectation.params != nil {
		mmCreateChat.mock.t.Fatalf("ChatServiceMock.CreateChat mock is already set by Expect")
	}

	if mmCreateChat.defaultExpectation.paramPtrs == nil {
		mmCreateChat.defaultExpectation.paramPtrs = &ChatServiceMockCreateChatParamPtrs{}
	}
	mmCreateChat.defaultExpectation.paramPtrs.req = &req
	mmCreateChat.defaultExpectation.expectationOrigins.originReq = minimock.CallerInfo(1)

	return mmCreateChat
}

// Inspect accepts an inspector function that has same arguments as the ChatService.CreateChat
func (mmCreateChat *mChatServiceMockCreateChat) Inspect(f func(ctx context.Context, req *chat_v1.CreateChatRequest)) *mChatServiceMockCreateChat {
	if mmCreateChat.mock.inspectFuncCreateChat != nil {
		mmCreateChat.mock.t.Fatalf("Inspect function is already set for ChatServiceMock.CreateChat")
	}

	mmCreateChat.mock.inspectFuncCreateChat = f

	return mmCreateChat
}

// Return sets up results that will be returned by ChatService.CreateChat
func (mmCreateChat *mChatServiceMockCreateChat) Return(i1 int64, err error) *ChatServiceMock {
	if mmCreateChat.mock.funcCreateChat != nil {
		mmCreateChat.mock.t.Fatalf("ChatServiceMock.CreateChat mock is already set by Set")
	}

	if mmCreateChat.defaultExpectation == nil {
		mmCreateChat.defaultExpectation = &ChatServiceMockCreateChatExpectation{mock: mmCreateChat.mock}
	}
	mmCreateChat.defaultExpectation.results = &ChatServiceMockCreateChatResults{i1, err}
	mmCreateChat.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmCreateChat.mock
}

// Set uses given function f to mock the ChatService.CreateChat method
func (mmCreateChat *mChatServiceMockCreateChat) Set(f func(ctx context.Context, req *chat_v1.CreateChatRequest) (i1 int64, err error)) *ChatServiceMock {
	if mmCreateChat.defaultExpectation != nil {
		mmCreateChat.mock.t.Fatalf("Default expectation is already set for the ChatService.CreateChat method")
	}

	if len(mmCreateChat.expectations) > 0 {
		mmCreateChat.mock.t.Fatalf("Some expectations are already set for the ChatService.CreateChat method")
	}

	mmCreateChat.mock.funcCreateChat = f
	mmCreateChat.mock.funcCreateChatOrigin = minimock.CallerInfo(1)
	return mmCreateChat.mock
}

// When sets expectation for the ChatService.CreateChat which will trigger the result defined by the following
// Then helper
func (mmCreateChat *mChatServiceMockCreateChat) When(ctx context.Context, req *chat_v1.CreateChatRequest) *ChatServiceMockCreateChatExpectation {
	if mmCreateChat.mock.funcCreateChat != nil {
		mmCreateChat.mock.t.Fatalf("ChatServiceMock.CreateChat mock is already set by Set")
	}

	expectation := &ChatServiceMockCreateChatExpectation{
		mock:               mmCreateChat.mock,
		params:             &ChatServiceMockCreateChatParams{ctx, req},
		expectationOrigins: ChatServiceMockCreateChatExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmCreateChat.expectations = append(mmCreateChat.expectations, expectation)
	return expectation
}

// Then sets up ChatService.CreateChat return parameters for the expectation previously defined by the When method
func (e *ChatServiceMockCreateChatExpectation) Then(i1 int64, err error) *ChatServiceMock {
	e.results = &ChatServiceMockCreateChatResults{i1, err}
	return e.mock
}

// Times sets number of times ChatService.CreateChat should be invoked
func (mmCreateChat *mChatServiceMockCreateChat) Times(n uint64) *mChatServiceMockCreateChat {
	if n == 0 {
		mmCreateChat.mock.t.Fatalf("Times of ChatServiceMock.CreateChat mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCreateChat.expectedInvocations, n)
	mmCreateChat.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmCreateChat
}

func (mmCreateChat *mChatServiceMockCreateChat) invocationsDone() bool {
	if len(mmCreateChat.expectations) == 0 && mmCreateChat.defaultExpectation == nil && mmCreateChat.mock.funcCreateChat == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCreateChat.mock.afterCreateChatCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCreateChat.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// CreateChat implements mm_service.ChatService
func (mmCreateChat *ChatServiceMock) CreateChat(ctx context.Context, req *chat_v1.CreateChatRequest) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreateChat.beforeCreateChatCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateChat.afterCreateChatCounter, 1)

	mmCreateChat.t.Helper()

	if mmCreateChat.inspectFuncCreateChat != nil {
		mmCreateChat.inspectFuncCreateChat(ctx, req)
	}

	mm_params := ChatServiceMockCreateChatParams{ctx, req}

	// Record call args
	mmCreateChat.CreateChatMock.mutex.Lock()
	mmCreateChat.CreateChatMock.callArgs = append(mmCreateChat.CreateChatMock.callArgs, &mm_params)
	mmCreateChat.CreateChatMock.mutex.Unlock()

	for _, e := range mmCreateChat.CreateChatMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmCreateChat.CreateChatMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateChat.CreateChatMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateChat.CreateChatMock.defaultExpectation.params
		mm_want_ptrs := mmCreateChat.CreateChatMock.defaultExpectation.paramPtrs

		mm_got := ChatServiceMockCreateChatParams{ctx, req}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmCreateChat.t.Errorf("ChatServiceMock.CreateChat got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreateChat.CreateChatMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.req != nil && !minimock.Equal(*mm_want_ptrs.req, mm_got.req) {
				mmCreateChat.t.Errorf("ChatServiceMock.CreateChat got unexpected parameter req, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreateChat.CreateChatMock.defaultExpectation.expectationOrigins.originReq, *mm_want_ptrs.req, mm_got.req, minimock.Diff(*mm_want_ptrs.req, mm_got.req))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateChat.t.Errorf("ChatServiceMock.CreateChat got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmCreateChat.CreateChatMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateChat.CreateChatMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateChat.t.Fatal("No results are set for the ChatServiceMock.CreateChat")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreateChat.funcCreateChat != nil {
		return mmCreateChat.funcCreateChat(ctx, req)
	}
	mmCreateChat.t.Fatalf("Unexpected call to ChatServiceMock.CreateChat. %v %v", ctx, req)
	return
}

// CreateChatAfterCounter returns a count of finished ChatServiceMock.CreateChat invocations
func (mmCreateChat *ChatServiceMock) CreateChatAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateChat.afterCreateChatCounter)
}

// CreateChatBeforeCounter returns a count of ChatServiceMock.CreateChat invocations
func (mmCreateChat *ChatServiceMock) CreateChatBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateChat.beforeCreateChatCounter)
}

// Calls returns a list of arguments used in each call to ChatServiceMock.CreateChat.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateChat *mChatServiceMockCreateChat) Calls() []*ChatServiceMockCreateChatParams {
	mmCreateChat.mutex.RLock()

	argCopy := make([]*ChatServiceMockCreateChatParams, len(mmCreateChat.callArgs))
	copy(argCopy, mmCreateChat.callArgs)

	mmCreateChat.mutex.RUnlock()

	return argCopy
}

// MinimockCreateChatDone returns true if the count of the CreateChat invocations corresponds
// the number of defined expectations
func (m *ChatServiceMock) MinimockCreateChatDone() bool {
	if m.CreateChatMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CreateChatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CreateChatMock.invocationsDone()
}

// MinimockCreateChatInspect logs each unmet expectation
func (m *ChatServiceMock) MinimockCreateChatInspect() {
	for _, e := range m.CreateChatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatServiceMock.CreateChat at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterCreateChatCounter := mm_atomic.LoadUint64(&m.afterCreateChatCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CreateChatMock.defaultExpectation != nil && afterCreateChatCounter < 1 {
		if m.CreateChatMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to ChatServiceMock.CreateChat at\n%s", m.CreateChatMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to ChatServiceMock.CreateChat at\n%s with params: %#v", m.CreateChatMock.defaultExpectation.expectationOrigins.origin, *m.CreateChatMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateChat != nil && afterCreateChatCounter < 1 {
		m.t.Errorf("Expected call to ChatServiceMock.CreateChat at\n%s", m.funcCreateChatOrigin)
	}

	if !m.CreateChatMock.invocationsDone() && afterCreateChatCounter > 0 {
		m.t.Errorf("Expected %d calls to ChatServiceMock.CreateChat at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.CreateChatMock.expectedInvocations), m.CreateChatMock.expectedInvocationsOrigin, afterCreateChatCounter)
	}
}

type mChatServiceMockDeleteChat struct {
	optional           bool
	mock               *ChatServiceMock
	defaultExpectation *ChatServiceMockDeleteChatExpectation
	expectations       []*ChatServiceMockDeleteChatExpectation

	callArgs []*ChatServiceMockDeleteChatParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// ChatServiceMockDeleteChatExpectation specifies expectation struct of the ChatService.DeleteChat
type ChatServiceMockDeleteChatExpectation struct {
	mock               *ChatServiceMock
	params             *ChatServiceMockDeleteChatParams
	paramPtrs          *ChatServiceMockDeleteChatParamPtrs
	expectationOrigins ChatServiceMockDeleteChatExpectationOrigins
	results            *ChatServiceMockDeleteChatResults
	returnOrigin       string
	Counter            uint64
}

// ChatServiceMockDeleteChatParams contains parameters of the ChatService.DeleteChat
type ChatServiceMockDeleteChatParams struct {
	ctx context.Context
	id  int64
}

// ChatServiceMockDeleteChatParamPtrs contains pointers to parameters of the ChatService.DeleteChat
type ChatServiceMockDeleteChatParamPtrs struct {
	ctx *context.Context
	id  *int64
}

// ChatServiceMockDeleteChatResults contains results of the ChatService.DeleteChat
type ChatServiceMockDeleteChatResults struct {
	err error
}

// ChatServiceMockDeleteChatOrigins contains origins of expectations of the ChatService.DeleteChat
type ChatServiceMockDeleteChatExpectationOrigins struct {
	origin    string
	originCtx string
	originId  string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmDeleteChat *mChatServiceMockDeleteChat) Optional() *mChatServiceMockDeleteChat {
	mmDeleteChat.optional = true
	return mmDeleteChat
}

// Expect sets up expected params for ChatService.DeleteChat
func (mmDeleteChat *mChatServiceMockDeleteChat) Expect(ctx context.Context, id int64) *mChatServiceMockDeleteChat {
	if mmDeleteChat.mock.funcDeleteChat != nil {
		mmDeleteChat.mock.t.Fatalf("ChatServiceMock.DeleteChat mock is already set by Set")
	}

	if mmDeleteChat.defaultExpectation == nil {
		mmDeleteChat.defaultExpectation = &ChatServiceMockDeleteChatExpectation{}
	}

	if mmDeleteChat.defaultExpectation.paramPtrs != nil {
		mmDeleteChat.mock.t.Fatalf("ChatServiceMock.DeleteChat mock is already set by ExpectParams functions")
	}

	mmDeleteChat.defaultExpectation.params = &ChatServiceMockDeleteChatParams{ctx, id}
	mmDeleteChat.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmDeleteChat.expectations {
		if minimock.Equal(e.params, mmDeleteChat.defaultExpectation.params) {
			mmDeleteChat.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDeleteChat.defaultExpectation.params)
		}
	}

	return mmDeleteChat
}

// ExpectCtxParam1 sets up expected param ctx for ChatService.DeleteChat
func (mmDeleteChat *mChatServiceMockDeleteChat) ExpectCtxParam1(ctx context.Context) *mChatServiceMockDeleteChat {
	if mmDeleteChat.mock.funcDeleteChat != nil {
		mmDeleteChat.mock.t.Fatalf("ChatServiceMock.DeleteChat mock is already set by Set")
	}

	if mmDeleteChat.defaultExpectation == nil {
		mmDeleteChat.defaultExpectation = &ChatServiceMockDeleteChatExpectation{}
	}

	if mmDeleteChat.defaultExpectation.params != nil {
		mmDeleteChat.mock.t.Fatalf("ChatServiceMock.DeleteChat mock is already set by Expect")
	}

	if mmDeleteChat.defaultExpectation.paramPtrs == nil {
		mmDeleteChat.defaultExpectation.paramPtrs = &ChatServiceMockDeleteChatParamPtrs{}
	}
	mmDeleteChat.defaultExpectation.paramPtrs.ctx = &ctx
	mmDeleteChat.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmDeleteChat
}

// ExpectIdParam2 sets up expected param id for ChatService.DeleteChat
func (mmDeleteChat *mChatServiceMockDeleteChat) ExpectIdParam2(id int64) *mChatServiceMockDeleteChat {
	if mmDeleteChat.mock.funcDeleteChat != nil {
		mmDeleteChat.mock.t.Fatalf("ChatServiceMock.DeleteChat mock is already set by Set")
	}

	if mmDeleteChat.defaultExpectation == nil {
		mmDeleteChat.defaultExpectation = &ChatServiceMockDeleteChatExpectation{}
	}

	if mmDeleteChat.defaultExpectation.params != nil {
		mmDeleteChat.mock.t.Fatalf("ChatServiceMock.DeleteChat mock is already set by Expect")
	}

	if mmDeleteChat.defaultExpectation.paramPtrs == nil {
		mmDeleteChat.defaultExpectation.paramPtrs = &ChatServiceMockDeleteChatParamPtrs{}
	}
	mmDeleteChat.defaultExpectation.paramPtrs.id = &id
	mmDeleteChat.defaultExpectation.expectationOrigins.originId = minimock.CallerInfo(1)

	return mmDeleteChat
}

// Inspect accepts an inspector function that has same arguments as the ChatService.DeleteChat
func (mmDeleteChat *mChatServiceMockDeleteChat) Inspect(f func(ctx context.Context, id int64)) *mChatServiceMockDeleteChat {
	if mmDeleteChat.mock.inspectFuncDeleteChat != nil {
		mmDeleteChat.mock.t.Fatalf("Inspect function is already set for ChatServiceMock.DeleteChat")
	}

	mmDeleteChat.mock.inspectFuncDeleteChat = f

	return mmDeleteChat
}

// Return sets up results that will be returned by ChatService.DeleteChat
func (mmDeleteChat *mChatServiceMockDeleteChat) Return(err error) *ChatServiceMock {
	if mmDeleteChat.mock.funcDeleteChat != nil {
		mmDeleteChat.mock.t.Fatalf("ChatServiceMock.DeleteChat mock is already set by Set")
	}

	if mmDeleteChat.defaultExpectation == nil {
		mmDeleteChat.defaultExpectation = &ChatServiceMockDeleteChatExpectation{mock: mmDeleteChat.mock}
	}
	mmDeleteChat.defaultExpectation.results = &ChatServiceMockDeleteChatResults{err}
	mmDeleteChat.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmDeleteChat.mock
}

// Set uses given function f to mock the ChatService.DeleteChat method
func (mmDeleteChat *mChatServiceMockDeleteChat) Set(f func(ctx context.Context, id int64) (err error)) *ChatServiceMock {
	if mmDeleteChat.defaultExpectation != nil {
		mmDeleteChat.mock.t.Fatalf("Default expectation is already set for the ChatService.DeleteChat method")
	}

	if len(mmDeleteChat.expectations) > 0 {
		mmDeleteChat.mock.t.Fatalf("Some expectations are already set for the ChatService.DeleteChat method")
	}

	mmDeleteChat.mock.funcDeleteChat = f
	mmDeleteChat.mock.funcDeleteChatOrigin = minimock.CallerInfo(1)
	return mmDeleteChat.mock
}

// When sets expectation for the ChatService.DeleteChat which will trigger the result defined by the following
// Then helper
func (mmDeleteChat *mChatServiceMockDeleteChat) When(ctx context.Context, id int64) *ChatServiceMockDeleteChatExpectation {
	if mmDeleteChat.mock.funcDeleteChat != nil {
		mmDeleteChat.mock.t.Fatalf("ChatServiceMock.DeleteChat mock is already set by Set")
	}

	expectation := &ChatServiceMockDeleteChatExpectation{
		mock:               mmDeleteChat.mock,
		params:             &ChatServiceMockDeleteChatParams{ctx, id},
		expectationOrigins: ChatServiceMockDeleteChatExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmDeleteChat.expectations = append(mmDeleteChat.expectations, expectation)
	return expectation
}

// Then sets up ChatService.DeleteChat return parameters for the expectation previously defined by the When method
func (e *ChatServiceMockDeleteChatExpectation) Then(err error) *ChatServiceMock {
	e.results = &ChatServiceMockDeleteChatResults{err}
	return e.mock
}

// Times sets number of times ChatService.DeleteChat should be invoked
func (mmDeleteChat *mChatServiceMockDeleteChat) Times(n uint64) *mChatServiceMockDeleteChat {
	if n == 0 {
		mmDeleteChat.mock.t.Fatalf("Times of ChatServiceMock.DeleteChat mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmDeleteChat.expectedInvocations, n)
	mmDeleteChat.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmDeleteChat
}

func (mmDeleteChat *mChatServiceMockDeleteChat) invocationsDone() bool {
	if len(mmDeleteChat.expectations) == 0 && mmDeleteChat.defaultExpectation == nil && mmDeleteChat.mock.funcDeleteChat == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmDeleteChat.mock.afterDeleteChatCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmDeleteChat.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// DeleteChat implements mm_service.ChatService
func (mmDeleteChat *ChatServiceMock) DeleteChat(ctx context.Context, id int64) (err error) {
	mm_atomic.AddUint64(&mmDeleteChat.beforeDeleteChatCounter, 1)
	defer mm_atomic.AddUint64(&mmDeleteChat.afterDeleteChatCounter, 1)

	mmDeleteChat.t.Helper()

	if mmDeleteChat.inspectFuncDeleteChat != nil {
		mmDeleteChat.inspectFuncDeleteChat(ctx, id)
	}

	mm_params := ChatServiceMockDeleteChatParams{ctx, id}

	// Record call args
	mmDeleteChat.DeleteChatMock.mutex.Lock()
	mmDeleteChat.DeleteChatMock.callArgs = append(mmDeleteChat.DeleteChatMock.callArgs, &mm_params)
	mmDeleteChat.DeleteChatMock.mutex.Unlock()

	for _, e := range mmDeleteChat.DeleteChatMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDeleteChat.DeleteChatMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDeleteChat.DeleteChatMock.defaultExpectation.Counter, 1)
		mm_want := mmDeleteChat.DeleteChatMock.defaultExpectation.params
		mm_want_ptrs := mmDeleteChat.DeleteChatMock.defaultExpectation.paramPtrs

		mm_got := ChatServiceMockDeleteChatParams{ctx, id}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmDeleteChat.t.Errorf("ChatServiceMock.DeleteChat got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmDeleteChat.DeleteChatMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.id != nil && !minimock.Equal(*mm_want_ptrs.id, mm_got.id) {
				mmDeleteChat.t.Errorf("ChatServiceMock.DeleteChat got unexpected parameter id, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmDeleteChat.DeleteChatMock.defaultExpectation.expectationOrigins.originId, *mm_want_ptrs.id, mm_got.id, minimock.Diff(*mm_want_ptrs.id, mm_got.id))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeleteChat.t.Errorf("ChatServiceMock.DeleteChat got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmDeleteChat.DeleteChatMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDeleteChat.DeleteChatMock.defaultExpectation.results
		if mm_results == nil {
			mmDeleteChat.t.Fatal("No results are set for the ChatServiceMock.DeleteChat")
		}
		return (*mm_results).err
	}
	if mmDeleteChat.funcDeleteChat != nil {
		return mmDeleteChat.funcDeleteChat(ctx, id)
	}
	mmDeleteChat.t.Fatalf("Unexpected call to ChatServiceMock.DeleteChat. %v %v", ctx, id)
	return
}

// DeleteChatAfterCounter returns a count of finished ChatServiceMock.DeleteChat invocations
func (mmDeleteChat *ChatServiceMock) DeleteChatAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteChat.afterDeleteChatCounter)
}

// DeleteChatBeforeCounter returns a count of ChatServiceMock.DeleteChat invocations
func (mmDeleteChat *ChatServiceMock) DeleteChatBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteChat.beforeDeleteChatCounter)
}

// Calls returns a list of arguments used in each call to ChatServiceMock.DeleteChat.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDeleteChat *mChatServiceMockDeleteChat) Calls() []*ChatServiceMockDeleteChatParams {
	mmDeleteChat.mutex.RLock()

	argCopy := make([]*ChatServiceMockDeleteChatParams, len(mmDeleteChat.callArgs))
	copy(argCopy, mmDeleteChat.callArgs)

	mmDeleteChat.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteChatDone returns true if the count of the DeleteChat invocations corresponds
// the number of defined expectations
func (m *ChatServiceMock) MinimockDeleteChatDone() bool {
	if m.DeleteChatMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.DeleteChatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.DeleteChatMock.invocationsDone()
}

// MinimockDeleteChatInspect logs each unmet expectation
func (m *ChatServiceMock) MinimockDeleteChatInspect() {
	for _, e := range m.DeleteChatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatServiceMock.DeleteChat at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterDeleteChatCounter := mm_atomic.LoadUint64(&m.afterDeleteChatCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteChatMock.defaultExpectation != nil && afterDeleteChatCounter < 1 {
		if m.DeleteChatMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to ChatServiceMock.DeleteChat at\n%s", m.DeleteChatMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to ChatServiceMock.DeleteChat at\n%s with params: %#v", m.DeleteChatMock.defaultExpectation.expectationOrigins.origin, *m.DeleteChatMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteChat != nil && afterDeleteChatCounter < 1 {
		m.t.Errorf("Expected call to ChatServiceMock.DeleteChat at\n%s", m.funcDeleteChatOrigin)
	}

	if !m.DeleteChatMock.invocationsDone() && afterDeleteChatCounter > 0 {
		m.t.Errorf("Expected %d calls to ChatServiceMock.DeleteChat at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.DeleteChatMock.expectedInvocations), m.DeleteChatMock.expectedInvocationsOrigin, afterDeleteChatCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ChatServiceMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateChatInspect()

			m.MinimockDeleteChatInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ChatServiceMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ChatServiceMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateChatDone() &&
		m.MinimockDeleteChatDone()
}
