package fluentd

import (
	"github.com/sirupsen/logrus"
)

type hookMock struct {
	instance *Hook
	CloseCount   uint
	LevelsCount  uint
	FireCount    uint
	FireArgs     []*logrus.Entry
	_Counts map[string]uint
	_GetQueueLengthCount uint
	_GetQueuedEntryAtArgs []uint
	_GetQueuedEntryAtCount uint
	_RemoveLogFromQueueArgs []uint
	_RemoveLogFromQueueCount uint
	_SendCount   uint
	_SendArgs    []map[string]interface{}
	_PostCount   uint
	_PostArgs    []hookMockPostArg
	_TraceArgs   []hookMockLogfArg
	_TraceCount  uint
	_DebugfCount uint
	_DebugfArgs  []hookMockLogfArg
	_WarnfCount  uint
	_WarnfArgs   []hookMockLogfArg
	_ErrorfCount uint
	_ErrorfArgs  []hookMockLogfArg
}

type hookMockLogfArg struct {
	format string
	args   []interface{}
}

type hookMockPostArg struct {
	level string
	data  map[string]interface{}
}

func (hook *hookMock) Levels() []logrus.Level {
	hook.LevelsCount += 1
	return []logrus.Level{}
}

func (hook *hookMock) Fire(entry *logrus.Entry) error {
	hook.FireArgs = append(hook.FireArgs, entry)
	hook.FireCount += 1
	return nil
}

func (hook *hookMock) Close() {
	hook.CloseCount += 1
}

func (hook *hookMock) getQueuedEntryAt(index uint) *logrus.Entry {
	hook._GetQueuedEntryAtArgs = append(hook._GetQueuedEntryAtArgs, index)
	hook._GetQueuedEntryAtCount += 1
	return hook.instance.getQueuedEntryAt(index)
}

func (hook *hookMock) getQueueLength() uint {
	hook._GetQueueLengthCount += 1
	return hook.instance.getQueueLength()
}

func (hook *hookMock) removeLogFromQueue(index uint) {
	hook.instance.removeLogFromQueue(index)
	hook._RemoveLogFromQueueArgs = append(hook._RemoveLogFromQueueArgs, index)
	hook._RemoveLogFromQueueCount += 1
}

func (hook *hookMock) shouldRetry() bool {
	hook._Counts["shouldRetry"] += 1
	return true
}

func (hook *hookMock) send(data map[string]interface{}) error {
	hook._SendCount += 1
	hook._SendArgs = append(hook._SendArgs, data)
	return nil
}

func (hook *hookMock) post(level string, data map[string]interface{}) {
	hook._PostArgs = append(hook._PostArgs, hookMockPostArg{level, data})
	hook._PostCount += 1
}

func (hook *hookMock) trace(message ...interface{}) {
	hook._TraceArgs = append(hook._TraceArgs, hookMockLogfArg{"", message})
	hook._TraceCount += 1
}

func (hook *hookMock) debugf(format string, others ...interface{}) {
	hook._DebugfArgs = append(hook._DebugfArgs, hookMockLogfArg{format, others})
	hook._DebugfCount += 1
}

func (hook *hookMock) warnf(format string, others ...interface{}) {
	hook._WarnfArgs = append(hook._WarnfArgs, hookMockLogfArg{format, others})
	hook._WarnfCount += 1
}

func (hook *hookMock) errorf(format string, others ...interface{}) {
	hook._ErrorfArgs = append(hook._ErrorfArgs, hookMockLogfArg{format, others})
	hook._ErrorfCount += 1
}
