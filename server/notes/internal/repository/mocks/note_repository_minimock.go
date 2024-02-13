// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package mocks

//go:generate minimock -i notes/internal/repository.NoteRepository -o note_repository_minimock.go -n NoteRepositoryMock -p mocks

import (
	"context"
	"notes/internal/model"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// NoteRepositoryMock implements repository.NoteRepository
type NoteRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreate          func(ctx context.Context, info *model.NoteInfo) (i1 int64, err error)
	inspectFuncCreate   func(ctx context.Context, info *model.NoteInfo)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mNoteRepositoryMockCreate

	funcGet          func(ctx context.Context, id int64) (np1 *model.Note, err error)
	inspectFuncGet   func(ctx context.Context, id int64)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mNoteRepositoryMockGet
}

// NewNoteRepositoryMock returns a mock for repository.NoteRepository
func NewNoteRepositoryMock(t minimock.Tester) *NoteRepositoryMock {
	m := &NoteRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mNoteRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*NoteRepositoryMockCreateParams{}

	m.GetMock = mNoteRepositoryMockGet{mock: m}
	m.GetMock.callArgs = []*NoteRepositoryMockGetParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mNoteRepositoryMockCreate struct {
	mock               *NoteRepositoryMock
	defaultExpectation *NoteRepositoryMockCreateExpectation
	expectations       []*NoteRepositoryMockCreateExpectation

	callArgs []*NoteRepositoryMockCreateParams
	mutex    sync.RWMutex
}

// NoteRepositoryMockCreateExpectation specifies expectation struct of the NoteRepository.Create
type NoteRepositoryMockCreateExpectation struct {
	mock    *NoteRepositoryMock
	params  *NoteRepositoryMockCreateParams
	results *NoteRepositoryMockCreateResults
	Counter uint64
}

// NoteRepositoryMockCreateParams contains parameters of the NoteRepository.Create
type NoteRepositoryMockCreateParams struct {
	ctx  context.Context
	info *model.NoteInfo
}

// NoteRepositoryMockCreateResults contains results of the NoteRepository.Create
type NoteRepositoryMockCreateResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for NoteRepository.Create
func (mmCreate *mNoteRepositoryMockCreate) Expect(ctx context.Context, info *model.NoteInfo) *mNoteRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("NoteRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &NoteRepositoryMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &NoteRepositoryMockCreateParams{ctx, info}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the NoteRepository.Create
func (mmCreate *mNoteRepositoryMockCreate) Inspect(f func(ctx context.Context, info *model.NoteInfo)) *mNoteRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for NoteRepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by NoteRepository.Create
func (mmCreate *mNoteRepositoryMockCreate) Return(i1 int64, err error) *NoteRepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("NoteRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &NoteRepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &NoteRepositoryMockCreateResults{i1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the NoteRepository.Create method
func (mmCreate *mNoteRepositoryMockCreate) Set(f func(ctx context.Context, info *model.NoteInfo) (i1 int64, err error)) *NoteRepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the NoteRepository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the NoteRepository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the NoteRepository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mNoteRepositoryMockCreate) When(ctx context.Context, info *model.NoteInfo) *NoteRepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("NoteRepositoryMock.Create mock is already set by Set")
	}

	expectation := &NoteRepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &NoteRepositoryMockCreateParams{ctx, info},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up NoteRepository.Create return parameters for the expectation previously defined by the When method
func (e *NoteRepositoryMockCreateExpectation) Then(i1 int64, err error) *NoteRepositoryMock {
	e.results = &NoteRepositoryMockCreateResults{i1, err}
	return e.mock
}

// Create implements repository.NoteRepository
func (mmCreate *NoteRepositoryMock) Create(ctx context.Context, info *model.NoteInfo) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, info)
	}

	mm_params := NoteRepositoryMockCreateParams{ctx, info}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, &mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := NoteRepositoryMockCreateParams{ctx, info}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("NoteRepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the NoteRepositoryMock.Create")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, info)
	}
	mmCreate.t.Fatalf("Unexpected call to NoteRepositoryMock.Create. %v %v", ctx, info)
	return
}

// CreateAfterCounter returns a count of finished NoteRepositoryMock.Create invocations
func (mmCreate *NoteRepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of NoteRepositoryMock.Create invocations
func (mmCreate *NoteRepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to NoteRepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mNoteRepositoryMockCreate) Calls() []*NoteRepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*NoteRepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *NoteRepositoryMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *NoteRepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to NoteRepositoryMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to NoteRepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to NoteRepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to NoteRepositoryMock.Create")
	}
}

type mNoteRepositoryMockGet struct {
	mock               *NoteRepositoryMock
	defaultExpectation *NoteRepositoryMockGetExpectation
	expectations       []*NoteRepositoryMockGetExpectation

	callArgs []*NoteRepositoryMockGetParams
	mutex    sync.RWMutex
}

// NoteRepositoryMockGetExpectation specifies expectation struct of the NoteRepository.Get
type NoteRepositoryMockGetExpectation struct {
	mock    *NoteRepositoryMock
	params  *NoteRepositoryMockGetParams
	results *NoteRepositoryMockGetResults
	Counter uint64
}

// NoteRepositoryMockGetParams contains parameters of the NoteRepository.Get
type NoteRepositoryMockGetParams struct {
	ctx context.Context
	id  int64
}

// NoteRepositoryMockGetResults contains results of the NoteRepository.Get
type NoteRepositoryMockGetResults struct {
	np1 *model.Note
	err error
}

// Expect sets up expected params for NoteRepository.Get
func (mmGet *mNoteRepositoryMockGet) Expect(ctx context.Context, id int64) *mNoteRepositoryMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("NoteRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &NoteRepositoryMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &NoteRepositoryMockGetParams{ctx, id}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the NoteRepository.Get
func (mmGet *mNoteRepositoryMockGet) Inspect(f func(ctx context.Context, id int64)) *mNoteRepositoryMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for NoteRepositoryMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by NoteRepository.Get
func (mmGet *mNoteRepositoryMockGet) Return(np1 *model.Note, err error) *NoteRepositoryMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("NoteRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &NoteRepositoryMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &NoteRepositoryMockGetResults{np1, err}
	return mmGet.mock
}

// Set uses given function f to mock the NoteRepository.Get method
func (mmGet *mNoteRepositoryMockGet) Set(f func(ctx context.Context, id int64) (np1 *model.Note, err error)) *NoteRepositoryMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the NoteRepository.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the NoteRepository.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the NoteRepository.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mNoteRepositoryMockGet) When(ctx context.Context, id int64) *NoteRepositoryMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("NoteRepositoryMock.Get mock is already set by Set")
	}

	expectation := &NoteRepositoryMockGetExpectation{
		mock:   mmGet.mock,
		params: &NoteRepositoryMockGetParams{ctx, id},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up NoteRepository.Get return parameters for the expectation previously defined by the When method
func (e *NoteRepositoryMockGetExpectation) Then(np1 *model.Note, err error) *NoteRepositoryMock {
	e.results = &NoteRepositoryMockGetResults{np1, err}
	return e.mock
}

// Get implements repository.NoteRepository
func (mmGet *NoteRepositoryMock) Get(ctx context.Context, id int64) (np1 *model.Note, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, id)
	}

	mm_params := NoteRepositoryMockGetParams{ctx, id}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, &mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.np1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := NoteRepositoryMockGetParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("NoteRepositoryMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the NoteRepositoryMock.Get")
		}
		return (*mm_results).np1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, id)
	}
	mmGet.t.Fatalf("Unexpected call to NoteRepositoryMock.Get. %v %v", ctx, id)
	return
}

// GetAfterCounter returns a count of finished NoteRepositoryMock.Get invocations
func (mmGet *NoteRepositoryMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of NoteRepositoryMock.Get invocations
func (mmGet *NoteRepositoryMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to NoteRepositoryMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mNoteRepositoryMockGet) Calls() []*NoteRepositoryMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*NoteRepositoryMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *NoteRepositoryMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *NoteRepositoryMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to NoteRepositoryMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to NoteRepositoryMock.Get")
		} else {
			m.t.Errorf("Expected call to NoteRepositoryMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to NoteRepositoryMock.Get")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *NoteRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateInspect()

			m.MinimockGetInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *NoteRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *NoteRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockGetDone()
}