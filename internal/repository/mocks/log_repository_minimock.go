// Code generated by http://github.com/gojuno/minimock (v3.4.1). DO NOT EDIT.

package mocks

//go:generate minimock -i github.com/valek177/chat-server/internal/repository.LogRepository -o log_repository_minimock.go -n LogRepositoryMock -p mocks

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/valek177/chat-server/internal/model"
)

// LogRepositoryMock implements mm_repository.LogRepository
type LogRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreateRecord          func(ctx context.Context, record *model.Record) (i1 int64, err error)
	funcCreateRecordOrigin    string
	inspectFuncCreateRecord   func(ctx context.Context, record *model.Record)
	afterCreateRecordCounter  uint64
	beforeCreateRecordCounter uint64
	CreateRecordMock          mLogRepositoryMockCreateRecord
}

// NewLogRepositoryMock returns a mock for mm_repository.LogRepository
func NewLogRepositoryMock(t minimock.Tester) *LogRepositoryMock {
	m := &LogRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateRecordMock = mLogRepositoryMockCreateRecord{mock: m}
	m.CreateRecordMock.callArgs = []*LogRepositoryMockCreateRecordParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mLogRepositoryMockCreateRecord struct {
	optional           bool
	mock               *LogRepositoryMock
	defaultExpectation *LogRepositoryMockCreateRecordExpectation
	expectations       []*LogRepositoryMockCreateRecordExpectation

	callArgs []*LogRepositoryMockCreateRecordParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// LogRepositoryMockCreateRecordExpectation specifies expectation struct of the LogRepository.CreateRecord
type LogRepositoryMockCreateRecordExpectation struct {
	mock               *LogRepositoryMock
	params             *LogRepositoryMockCreateRecordParams
	paramPtrs          *LogRepositoryMockCreateRecordParamPtrs
	expectationOrigins LogRepositoryMockCreateRecordExpectationOrigins
	results            *LogRepositoryMockCreateRecordResults
	returnOrigin       string
	Counter            uint64
}

// LogRepositoryMockCreateRecordParams contains parameters of the LogRepository.CreateRecord
type LogRepositoryMockCreateRecordParams struct {
	ctx    context.Context
	record *model.Record
}

// LogRepositoryMockCreateRecordParamPtrs contains pointers to parameters of the LogRepository.CreateRecord
type LogRepositoryMockCreateRecordParamPtrs struct {
	ctx    *context.Context
	record **model.Record
}

// LogRepositoryMockCreateRecordResults contains results of the LogRepository.CreateRecord
type LogRepositoryMockCreateRecordResults struct {
	i1  int64
	err error
}

// LogRepositoryMockCreateRecordOrigins contains origins of expectations of the LogRepository.CreateRecord
type LogRepositoryMockCreateRecordExpectationOrigins struct {
	origin       string
	originCtx    string
	originRecord string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmCreateRecord *mLogRepositoryMockCreateRecord) Optional() *mLogRepositoryMockCreateRecord {
	mmCreateRecord.optional = true
	return mmCreateRecord
}

// Expect sets up expected params for LogRepository.CreateRecord
func (mmCreateRecord *mLogRepositoryMockCreateRecord) Expect(ctx context.Context, record *model.Record) *mLogRepositoryMockCreateRecord {
	if mmCreateRecord.mock.funcCreateRecord != nil {
		mmCreateRecord.mock.t.Fatalf("LogRepositoryMock.CreateRecord mock is already set by Set")
	}

	if mmCreateRecord.defaultExpectation == nil {
		mmCreateRecord.defaultExpectation = &LogRepositoryMockCreateRecordExpectation{}
	}

	if mmCreateRecord.defaultExpectation.paramPtrs != nil {
		mmCreateRecord.mock.t.Fatalf("LogRepositoryMock.CreateRecord mock is already set by ExpectParams functions")
	}

	mmCreateRecord.defaultExpectation.params = &LogRepositoryMockCreateRecordParams{ctx, record}
	mmCreateRecord.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmCreateRecord.expectations {
		if minimock.Equal(e.params, mmCreateRecord.defaultExpectation.params) {
			mmCreateRecord.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreateRecord.defaultExpectation.params)
		}
	}

	return mmCreateRecord
}

// ExpectCtxParam1 sets up expected param ctx for LogRepository.CreateRecord
func (mmCreateRecord *mLogRepositoryMockCreateRecord) ExpectCtxParam1(ctx context.Context) *mLogRepositoryMockCreateRecord {
	if mmCreateRecord.mock.funcCreateRecord != nil {
		mmCreateRecord.mock.t.Fatalf("LogRepositoryMock.CreateRecord mock is already set by Set")
	}

	if mmCreateRecord.defaultExpectation == nil {
		mmCreateRecord.defaultExpectation = &LogRepositoryMockCreateRecordExpectation{}
	}

	if mmCreateRecord.defaultExpectation.params != nil {
		mmCreateRecord.mock.t.Fatalf("LogRepositoryMock.CreateRecord mock is already set by Expect")
	}

	if mmCreateRecord.defaultExpectation.paramPtrs == nil {
		mmCreateRecord.defaultExpectation.paramPtrs = &LogRepositoryMockCreateRecordParamPtrs{}
	}
	mmCreateRecord.defaultExpectation.paramPtrs.ctx = &ctx
	mmCreateRecord.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmCreateRecord
}

// ExpectRecordParam2 sets up expected param record for LogRepository.CreateRecord
func (mmCreateRecord *mLogRepositoryMockCreateRecord) ExpectRecordParam2(record *model.Record) *mLogRepositoryMockCreateRecord {
	if mmCreateRecord.mock.funcCreateRecord != nil {
		mmCreateRecord.mock.t.Fatalf("LogRepositoryMock.CreateRecord mock is already set by Set")
	}

	if mmCreateRecord.defaultExpectation == nil {
		mmCreateRecord.defaultExpectation = &LogRepositoryMockCreateRecordExpectation{}
	}

	if mmCreateRecord.defaultExpectation.params != nil {
		mmCreateRecord.mock.t.Fatalf("LogRepositoryMock.CreateRecord mock is already set by Expect")
	}

	if mmCreateRecord.defaultExpectation.paramPtrs == nil {
		mmCreateRecord.defaultExpectation.paramPtrs = &LogRepositoryMockCreateRecordParamPtrs{}
	}
	mmCreateRecord.defaultExpectation.paramPtrs.record = &record
	mmCreateRecord.defaultExpectation.expectationOrigins.originRecord = minimock.CallerInfo(1)

	return mmCreateRecord
}

// Inspect accepts an inspector function that has same arguments as the LogRepository.CreateRecord
func (mmCreateRecord *mLogRepositoryMockCreateRecord) Inspect(f func(ctx context.Context, record *model.Record)) *mLogRepositoryMockCreateRecord {
	if mmCreateRecord.mock.inspectFuncCreateRecord != nil {
		mmCreateRecord.mock.t.Fatalf("Inspect function is already set for LogRepositoryMock.CreateRecord")
	}

	mmCreateRecord.mock.inspectFuncCreateRecord = f

	return mmCreateRecord
}

// Return sets up results that will be returned by LogRepository.CreateRecord
func (mmCreateRecord *mLogRepositoryMockCreateRecord) Return(i1 int64, err error) *LogRepositoryMock {
	if mmCreateRecord.mock.funcCreateRecord != nil {
		mmCreateRecord.mock.t.Fatalf("LogRepositoryMock.CreateRecord mock is already set by Set")
	}

	if mmCreateRecord.defaultExpectation == nil {
		mmCreateRecord.defaultExpectation = &LogRepositoryMockCreateRecordExpectation{mock: mmCreateRecord.mock}
	}
	mmCreateRecord.defaultExpectation.results = &LogRepositoryMockCreateRecordResults{i1, err}
	mmCreateRecord.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmCreateRecord.mock
}

// Set uses given function f to mock the LogRepository.CreateRecord method
func (mmCreateRecord *mLogRepositoryMockCreateRecord) Set(f func(ctx context.Context, record *model.Record) (i1 int64, err error)) *LogRepositoryMock {
	if mmCreateRecord.defaultExpectation != nil {
		mmCreateRecord.mock.t.Fatalf("Default expectation is already set for the LogRepository.CreateRecord method")
	}

	if len(mmCreateRecord.expectations) > 0 {
		mmCreateRecord.mock.t.Fatalf("Some expectations are already set for the LogRepository.CreateRecord method")
	}

	mmCreateRecord.mock.funcCreateRecord = f
	mmCreateRecord.mock.funcCreateRecordOrigin = minimock.CallerInfo(1)
	return mmCreateRecord.mock
}

// When sets expectation for the LogRepository.CreateRecord which will trigger the result defined by the following
// Then helper
func (mmCreateRecord *mLogRepositoryMockCreateRecord) When(ctx context.Context, record *model.Record) *LogRepositoryMockCreateRecordExpectation {
	if mmCreateRecord.mock.funcCreateRecord != nil {
		mmCreateRecord.mock.t.Fatalf("LogRepositoryMock.CreateRecord mock is already set by Set")
	}

	expectation := &LogRepositoryMockCreateRecordExpectation{
		mock:               mmCreateRecord.mock,
		params:             &LogRepositoryMockCreateRecordParams{ctx, record},
		expectationOrigins: LogRepositoryMockCreateRecordExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmCreateRecord.expectations = append(mmCreateRecord.expectations, expectation)
	return expectation
}

// Then sets up LogRepository.CreateRecord return parameters for the expectation previously defined by the When method
func (e *LogRepositoryMockCreateRecordExpectation) Then(i1 int64, err error) *LogRepositoryMock {
	e.results = &LogRepositoryMockCreateRecordResults{i1, err}
	return e.mock
}

// Times sets number of times LogRepository.CreateRecord should be invoked
func (mmCreateRecord *mLogRepositoryMockCreateRecord) Times(n uint64) *mLogRepositoryMockCreateRecord {
	if n == 0 {
		mmCreateRecord.mock.t.Fatalf("Times of LogRepositoryMock.CreateRecord mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmCreateRecord.expectedInvocations, n)
	mmCreateRecord.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmCreateRecord
}

func (mmCreateRecord *mLogRepositoryMockCreateRecord) invocationsDone() bool {
	if len(mmCreateRecord.expectations) == 0 && mmCreateRecord.defaultExpectation == nil && mmCreateRecord.mock.funcCreateRecord == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmCreateRecord.mock.afterCreateRecordCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmCreateRecord.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// CreateRecord implements mm_repository.LogRepository
func (mmCreateRecord *LogRepositoryMock) CreateRecord(ctx context.Context, record *model.Record) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreateRecord.beforeCreateRecordCounter, 1)
	defer mm_atomic.AddUint64(&mmCreateRecord.afterCreateRecordCounter, 1)

	mmCreateRecord.t.Helper()

	if mmCreateRecord.inspectFuncCreateRecord != nil {
		mmCreateRecord.inspectFuncCreateRecord(ctx, record)
	}

	mm_params := LogRepositoryMockCreateRecordParams{ctx, record}

	// Record call args
	mmCreateRecord.CreateRecordMock.mutex.Lock()
	mmCreateRecord.CreateRecordMock.callArgs = append(mmCreateRecord.CreateRecordMock.callArgs, &mm_params)
	mmCreateRecord.CreateRecordMock.mutex.Unlock()

	for _, e := range mmCreateRecord.CreateRecordMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmCreateRecord.CreateRecordMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreateRecord.CreateRecordMock.defaultExpectation.Counter, 1)
		mm_want := mmCreateRecord.CreateRecordMock.defaultExpectation.params
		mm_want_ptrs := mmCreateRecord.CreateRecordMock.defaultExpectation.paramPtrs

		mm_got := LogRepositoryMockCreateRecordParams{ctx, record}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmCreateRecord.t.Errorf("LogRepositoryMock.CreateRecord got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreateRecord.CreateRecordMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.record != nil && !minimock.Equal(*mm_want_ptrs.record, mm_got.record) {
				mmCreateRecord.t.Errorf("LogRepositoryMock.CreateRecord got unexpected parameter record, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmCreateRecord.CreateRecordMock.defaultExpectation.expectationOrigins.originRecord, *mm_want_ptrs.record, mm_got.record, minimock.Diff(*mm_want_ptrs.record, mm_got.record))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreateRecord.t.Errorf("LogRepositoryMock.CreateRecord got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmCreateRecord.CreateRecordMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreateRecord.CreateRecordMock.defaultExpectation.results
		if mm_results == nil {
			mmCreateRecord.t.Fatal("No results are set for the LogRepositoryMock.CreateRecord")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreateRecord.funcCreateRecord != nil {
		return mmCreateRecord.funcCreateRecord(ctx, record)
	}
	mmCreateRecord.t.Fatalf("Unexpected call to LogRepositoryMock.CreateRecord. %v %v", ctx, record)
	return
}

// CreateRecordAfterCounter returns a count of finished LogRepositoryMock.CreateRecord invocations
func (mmCreateRecord *LogRepositoryMock) CreateRecordAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateRecord.afterCreateRecordCounter)
}

// CreateRecordBeforeCounter returns a count of LogRepositoryMock.CreateRecord invocations
func (mmCreateRecord *LogRepositoryMock) CreateRecordBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreateRecord.beforeCreateRecordCounter)
}

// Calls returns a list of arguments used in each call to LogRepositoryMock.CreateRecord.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreateRecord *mLogRepositoryMockCreateRecord) Calls() []*LogRepositoryMockCreateRecordParams {
	mmCreateRecord.mutex.RLock()

	argCopy := make([]*LogRepositoryMockCreateRecordParams, len(mmCreateRecord.callArgs))
	copy(argCopy, mmCreateRecord.callArgs)

	mmCreateRecord.mutex.RUnlock()

	return argCopy
}

// MinimockCreateRecordDone returns true if the count of the CreateRecord invocations corresponds
// the number of defined expectations
func (m *LogRepositoryMock) MinimockCreateRecordDone() bool {
	if m.CreateRecordMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.CreateRecordMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.CreateRecordMock.invocationsDone()
}

// MinimockCreateRecordInspect logs each unmet expectation
func (m *LogRepositoryMock) MinimockCreateRecordInspect() {
	for _, e := range m.CreateRecordMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to LogRepositoryMock.CreateRecord at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterCreateRecordCounter := mm_atomic.LoadUint64(&m.afterCreateRecordCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.CreateRecordMock.defaultExpectation != nil && afterCreateRecordCounter < 1 {
		if m.CreateRecordMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to LogRepositoryMock.CreateRecord at\n%s", m.CreateRecordMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to LogRepositoryMock.CreateRecord at\n%s with params: %#v", m.CreateRecordMock.defaultExpectation.expectationOrigins.origin, *m.CreateRecordMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreateRecord != nil && afterCreateRecordCounter < 1 {
		m.t.Errorf("Expected call to LogRepositoryMock.CreateRecord at\n%s", m.funcCreateRecordOrigin)
	}

	if !m.CreateRecordMock.invocationsDone() && afterCreateRecordCounter > 0 {
		m.t.Errorf("Expected %d calls to LogRepositoryMock.CreateRecord at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.CreateRecordMock.expectedInvocations), m.CreateRecordMock.expectedInvocationsOrigin, afterCreateRecordCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *LogRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateRecordInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *LogRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *LogRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateRecordDone()
}