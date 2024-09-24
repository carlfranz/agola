// Code generated by go generate; DO NOT EDIT.
package db

import (
	"encoding/json"
	stdsql "database/sql"
	"time"

	"github.com/sorintlab/errors"
	sq "github.com/huandu/go-sqlbuilder"

	"agola.io/agola/internal/sqlg/sql"

	types "agola.io/agola/services/runservice/types"
)
var (
	changeGroupInsertPostgres = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inName string, inValue string) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("changegroup").Cols("id", "revision", "creation_time", "update_time", "name", "value").Values(inID, inRevision, inCreationTime, inUpdateTime, inName, inValue)
	}
	changeGroupUpdatePostgres = func(curRevision uint64, inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inName string, inValue string) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("changegroup").Set(ub.Assign("id", inID), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("name", inName), ub.Assign("value", inValue)).Where(ub.E("id", inID), ub.E("revision", curRevision))
	}

	changeGroupInsertRawPostgres = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inName string, inValue string) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("changegroup").Cols("id", "revision", "creation_time", "update_time", "name", "value").SQL("OVERRIDING SYSTEM VALUE").Values(inID, inRevision, inCreationTime, inUpdateTime, inName, inValue)
	}
)

func (d *DB) insertChangeGroupPostgres(tx *sql.Tx, changegroup *types.ChangeGroup) error {
	q := changeGroupInsertPostgres(changegroup.ID, changegroup.Revision, changegroup.CreationTime, changegroup.UpdateTime, changegroup.Name, changegroup.Value)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert changeGroup")
	}

	return nil
}

func (d *DB) updateChangeGroupPostgres(tx *sql.Tx, curRevision uint64, changegroup *types.ChangeGroup) (stdsql.Result, error) {
	q := changeGroupUpdatePostgres(curRevision, changegroup.ID, changegroup.Revision, changegroup.CreationTime, changegroup.UpdateTime, changegroup.Name, changegroup.Value)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update changeGroup")
	}

	return res, nil
}

func (d *DB) insertRawChangeGroupPostgres(tx *sql.Tx, changegroup *types.ChangeGroup) error {
	q := changeGroupInsertRawPostgres(changegroup.ID, changegroup.Revision, changegroup.CreationTime, changegroup.UpdateTime, changegroup.Name, changegroup.Value)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert changeGroup")
	}

	return nil
}
var (
	runConfigInsertPostgres = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inName string, inGroup string, inSetupErrors []byte, inAnnotations []byte, inStaticEnvironment []byte, inEnvironment []byte, inTasks []byte, inCacheGroup string) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("runconfig").Cols("id", "revision", "creation_time", "update_time", "name", "run_group", "setup_errors", "annotations", "static_environment", "environment", "tasks", "cache_group").Values(inID, inRevision, inCreationTime, inUpdateTime, inName, inGroup, inSetupErrors, inAnnotations, inStaticEnvironment, inEnvironment, inTasks, inCacheGroup)
	}
	runConfigUpdatePostgres = func(curRevision uint64, inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inName string, inGroup string, inSetupErrors []byte, inAnnotations []byte, inStaticEnvironment []byte, inEnvironment []byte, inTasks []byte, inCacheGroup string) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("runconfig").Set(ub.Assign("id", inID), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("name", inName), ub.Assign("run_group", inGroup), ub.Assign("setup_errors", inSetupErrors), ub.Assign("annotations", inAnnotations), ub.Assign("static_environment", inStaticEnvironment), ub.Assign("environment", inEnvironment), ub.Assign("tasks", inTasks), ub.Assign("cache_group", inCacheGroup)).Where(ub.E("id", inID), ub.E("revision", curRevision))
	}

	runConfigInsertRawPostgres = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inName string, inGroup string, inSetupErrors []byte, inAnnotations []byte, inStaticEnvironment []byte, inEnvironment []byte, inTasks []byte, inCacheGroup string) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("runconfig").Cols("id", "revision", "creation_time", "update_time", "name", "run_group", "setup_errors", "annotations", "static_environment", "environment", "tasks", "cache_group").SQL("OVERRIDING SYSTEM VALUE").Values(inID, inRevision, inCreationTime, inUpdateTime, inName, inGroup, inSetupErrors, inAnnotations, inStaticEnvironment, inEnvironment, inTasks, inCacheGroup)
	}
)

func (d *DB) insertRunConfigPostgres(tx *sql.Tx, runconfig *types.RunConfig) error {
	inSetupErrorsJSON, err := json.Marshal(runconfig.SetupErrors)
	if err != nil {
		return errors.Wrap(err, "failed to marshal runconfig.SetupErrors")
	}
	inAnnotationsJSON, err := json.Marshal(runconfig.Annotations)
	if err != nil {
		return errors.Wrap(err, "failed to marshal runconfig.Annotations")
	}
	inStaticEnvironmentJSON, err := json.Marshal(runconfig.StaticEnvironment)
	if err != nil {
		return errors.Wrap(err, "failed to marshal runconfig.StaticEnvironment")
	}
	inEnvironmentJSON, err := json.Marshal(runconfig.Environment)
	if err != nil {
		return errors.Wrap(err, "failed to marshal runconfig.Environment")
	}
	inTasksJSON, err := json.Marshal(runconfig.Tasks)
	if err != nil {
		return errors.Wrap(err, "failed to marshal runconfig.Tasks")
	}
	q := runConfigInsertPostgres(runconfig.ID, runconfig.Revision, runconfig.CreationTime, runconfig.UpdateTime, runconfig.Name, runconfig.Group, inSetupErrorsJSON, inAnnotationsJSON, inStaticEnvironmentJSON, inEnvironmentJSON, inTasksJSON, runconfig.CacheGroup)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert runConfig")
	}

	return nil
}

func (d *DB) updateRunConfigPostgres(tx *sql.Tx, curRevision uint64, runconfig *types.RunConfig) (stdsql.Result, error) {
	inSetupErrorsJSON, err := json.Marshal(runconfig.SetupErrors)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal runconfig.SetupErrors")
	}
	inAnnotationsJSON, err := json.Marshal(runconfig.Annotations)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal runconfig.Annotations")
	}
	inStaticEnvironmentJSON, err := json.Marshal(runconfig.StaticEnvironment)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal runconfig.StaticEnvironment")
	}
	inEnvironmentJSON, err := json.Marshal(runconfig.Environment)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal runconfig.Environment")
	}
	inTasksJSON, err := json.Marshal(runconfig.Tasks)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal runconfig.Tasks")
	}
	q := runConfigUpdatePostgres(curRevision, runconfig.ID, runconfig.Revision, runconfig.CreationTime, runconfig.UpdateTime, runconfig.Name, runconfig.Group, inSetupErrorsJSON, inAnnotationsJSON, inStaticEnvironmentJSON, inEnvironmentJSON, inTasksJSON, runconfig.CacheGroup)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update runConfig")
	}

	return res, nil
}

func (d *DB) insertRawRunConfigPostgres(tx *sql.Tx, runconfig *types.RunConfig) error {
	inSetupErrorsJSON, err := json.Marshal(runconfig.SetupErrors)
	if err != nil {
		return errors.Wrap(err, "failed to marshal runconfig.SetupErrors")
	}
	inAnnotationsJSON, err := json.Marshal(runconfig.Annotations)
	if err != nil {
		return errors.Wrap(err, "failed to marshal runconfig.Annotations")
	}
	inStaticEnvironmentJSON, err := json.Marshal(runconfig.StaticEnvironment)
	if err != nil {
		return errors.Wrap(err, "failed to marshal runconfig.StaticEnvironment")
	}
	inEnvironmentJSON, err := json.Marshal(runconfig.Environment)
	if err != nil {
		return errors.Wrap(err, "failed to marshal runconfig.Environment")
	}
	inTasksJSON, err := json.Marshal(runconfig.Tasks)
	if err != nil {
		return errors.Wrap(err, "failed to marshal runconfig.Tasks")
	}
	q := runConfigInsertRawPostgres(runconfig.ID, runconfig.Revision, runconfig.CreationTime, runconfig.UpdateTime, runconfig.Name, runconfig.Group, inSetupErrorsJSON, inAnnotationsJSON, inStaticEnvironmentJSON, inEnvironmentJSON, inTasksJSON, runconfig.CacheGroup)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert runConfig")
	}

	return nil
}
var (
	runInsertPostgres = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inName string, inRunConfigID string, inCounter uint64, inGroup string, inAnnotations []byte, inPhase types.RunPhase, inResult types.RunResult, inStop bool, inTasks []byte, inEnqueueTime *time.Time, inStartTime *time.Time, inEndTime *time.Time, inArchived bool) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("run").Cols("id", "revision", "creation_time", "update_time", "name", "run_config_id", "counter", "run_group", "annotations", "phase", "result", "stop", "tasks", "enqueue_time", "start_time", "end_time", "archived").Values(inID, inRevision, inCreationTime, inUpdateTime, inName, inRunConfigID, inCounter, inGroup, inAnnotations, inPhase, inResult, inStop, inTasks, inEnqueueTime, inStartTime, inEndTime, inArchived)
	}
	runUpdatePostgres = func(curRevision uint64, inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inName string, inRunConfigID string, inCounter uint64, inGroup string, inAnnotations []byte, inPhase types.RunPhase, inResult types.RunResult, inStop bool, inTasks []byte, inEnqueueTime *time.Time, inStartTime *time.Time, inEndTime *time.Time, inArchived bool) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("run").Set(ub.Assign("id", inID), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("name", inName), ub.Assign("run_config_id", inRunConfigID), ub.Assign("counter", inCounter), ub.Assign("run_group", inGroup), ub.Assign("annotations", inAnnotations), ub.Assign("phase", inPhase), ub.Assign("result", inResult), ub.Assign("stop", inStop), ub.Assign("tasks", inTasks), ub.Assign("enqueue_time", inEnqueueTime), ub.Assign("start_time", inStartTime), ub.Assign("end_time", inEndTime), ub.Assign("archived", inArchived)).Where(ub.E("id", inID), ub.E("revision", curRevision))
	}

	runInsertRawPostgres = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inSequence uint64, inName string, inRunConfigID string, inCounter uint64, inGroup string, inAnnotations []byte, inPhase types.RunPhase, inResult types.RunResult, inStop bool, inTasks []byte, inEnqueueTime *time.Time, inStartTime *time.Time, inEndTime *time.Time, inArchived bool) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("run").Cols("id", "revision", "creation_time", "update_time", "sequence", "name", "run_config_id", "counter", "run_group", "annotations", "phase", "result", "stop", "tasks", "enqueue_time", "start_time", "end_time", "archived").SQL("OVERRIDING SYSTEM VALUE").Values(inID, inRevision, inCreationTime, inUpdateTime, inSequence, inName, inRunConfigID, inCounter, inGroup, inAnnotations, inPhase, inResult, inStop, inTasks, inEnqueueTime, inStartTime, inEndTime, inArchived)
	}
)

func (d *DB) insertRunPostgres(tx *sql.Tx, run *types.Run) error {
	inAnnotationsJSON, err := json.Marshal(run.Annotations)
	if err != nil {
		return errors.Wrap(err, "failed to marshal run.Annotations")
	}
	inTasksJSON, err := json.Marshal(run.Tasks)
	if err != nil {
		return errors.Wrap(err, "failed to marshal run.Tasks")
	}
	q := runInsertPostgres(run.ID, run.Revision, run.CreationTime, run.UpdateTime, run.Name, run.RunConfigID, run.Counter, run.Group, inAnnotationsJSON, run.Phase, run.Result, run.Stop, inTasksJSON, run.EnqueueTime, run.StartTime, run.EndTime, run.Archived)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert run")
	}

	return nil
}

func (d *DB) updateRunPostgres(tx *sql.Tx, curRevision uint64, run *types.Run) (stdsql.Result, error) {
	inAnnotationsJSON, err := json.Marshal(run.Annotations)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal run.Annotations")
	}
	inTasksJSON, err := json.Marshal(run.Tasks)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal run.Tasks")
	}
	q := runUpdatePostgres(curRevision, run.ID, run.Revision, run.CreationTime, run.UpdateTime, run.Name, run.RunConfigID, run.Counter, run.Group, inAnnotationsJSON, run.Phase, run.Result, run.Stop, inTasksJSON, run.EnqueueTime, run.StartTime, run.EndTime, run.Archived)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update run")
	}

	return res, nil
}

func (d *DB) insertRawRunPostgres(tx *sql.Tx, run *types.Run) error {
	inAnnotationsJSON, err := json.Marshal(run.Annotations)
	if err != nil {
		return errors.Wrap(err, "failed to marshal run.Annotations")
	}
	inTasksJSON, err := json.Marshal(run.Tasks)
	if err != nil {
		return errors.Wrap(err, "failed to marshal run.Tasks")
	}
	q := runInsertRawPostgres(run.ID, run.Revision, run.CreationTime, run.UpdateTime, run.Sequence, run.Name, run.RunConfigID, run.Counter, run.Group, inAnnotationsJSON, run.Phase, run.Result, run.Stop, inTasksJSON, run.EnqueueTime, run.StartTime, run.EndTime, run.Archived)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert run")
	}

	return nil
}
var (
	runCounterInsertPostgres = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inGroupID string, inValue uint64) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("runcounter").Cols("id", "revision", "creation_time", "update_time", "group_id", "value").Values(inID, inRevision, inCreationTime, inUpdateTime, inGroupID, inValue)
	}
	runCounterUpdatePostgres = func(curRevision uint64, inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inGroupID string, inValue uint64) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("runcounter").Set(ub.Assign("id", inID), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("group_id", inGroupID), ub.Assign("value", inValue)).Where(ub.E("id", inID), ub.E("revision", curRevision))
	}

	runCounterInsertRawPostgres = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inGroupID string, inValue uint64) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("runcounter").Cols("id", "revision", "creation_time", "update_time", "group_id", "value").SQL("OVERRIDING SYSTEM VALUE").Values(inID, inRevision, inCreationTime, inUpdateTime, inGroupID, inValue)
	}
)

func (d *DB) insertRunCounterPostgres(tx *sql.Tx, runcounter *types.RunCounter) error {
	q := runCounterInsertPostgres(runcounter.ID, runcounter.Revision, runcounter.CreationTime, runcounter.UpdateTime, runcounter.GroupID, runcounter.Value)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert runCounter")
	}

	return nil
}

func (d *DB) updateRunCounterPostgres(tx *sql.Tx, curRevision uint64, runcounter *types.RunCounter) (stdsql.Result, error) {
	q := runCounterUpdatePostgres(curRevision, runcounter.ID, runcounter.Revision, runcounter.CreationTime, runcounter.UpdateTime, runcounter.GroupID, runcounter.Value)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update runCounter")
	}

	return res, nil
}

func (d *DB) insertRawRunCounterPostgres(tx *sql.Tx, runcounter *types.RunCounter) error {
	q := runCounterInsertRawPostgres(runcounter.ID, runcounter.Revision, runcounter.CreationTime, runcounter.UpdateTime, runcounter.GroupID, runcounter.Value)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert runCounter")
	}

	return nil
}
var (
	runEventInsertPostgres = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inRunEventType types.RunEventType, inRunID string, inPhase types.RunPhase, inResult types.RunResult, inData []byte, inDataVersion uint64) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("runevent").Cols("id", "revision", "creation_time", "update_time", "run_event_type", "run_id", "phase", "result", "data", "data_version").Values(inID, inRevision, inCreationTime, inUpdateTime, inRunEventType, inRunID, inPhase, inResult, inData, inDataVersion)
	}
	runEventUpdatePostgres = func(curRevision uint64, inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inRunEventType types.RunEventType, inRunID string, inPhase types.RunPhase, inResult types.RunResult, inData []byte, inDataVersion uint64) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("runevent").Set(ub.Assign("id", inID), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("run_event_type", inRunEventType), ub.Assign("run_id", inRunID), ub.Assign("phase", inPhase), ub.Assign("result", inResult), ub.Assign("data", inData), ub.Assign("data_version", inDataVersion)).Where(ub.E("id", inID), ub.E("revision", curRevision))
	}

	runEventInsertRawPostgres = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inSequence uint64, inRunEventType types.RunEventType, inRunID string, inPhase types.RunPhase, inResult types.RunResult, inData []byte, inDataVersion uint64) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("runevent").Cols("id", "revision", "creation_time", "update_time", "sequence", "run_event_type", "run_id", "phase", "result", "data", "data_version").SQL("OVERRIDING SYSTEM VALUE").Values(inID, inRevision, inCreationTime, inUpdateTime, inSequence, inRunEventType, inRunID, inPhase, inResult, inData, inDataVersion)
	}
)

func (d *DB) insertRunEventPostgres(tx *sql.Tx, runevent *types.RunEvent) error {
	inDataJSON, err := json.Marshal(runevent.Data)
	if err != nil {
		return errors.Wrap(err, "failed to marshal runevent.Data")
	}
	q := runEventInsertPostgres(runevent.ID, runevent.Revision, runevent.CreationTime, runevent.UpdateTime, runevent.RunEventType, runevent.RunID, runevent.Phase, runevent.Result, inDataJSON, runevent.DataVersion)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert runEvent")
	}

	return nil
}

func (d *DB) updateRunEventPostgres(tx *sql.Tx, curRevision uint64, runevent *types.RunEvent) (stdsql.Result, error) {
	inDataJSON, err := json.Marshal(runevent.Data)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal runevent.Data")
	}
	q := runEventUpdatePostgres(curRevision, runevent.ID, runevent.Revision, runevent.CreationTime, runevent.UpdateTime, runevent.RunEventType, runevent.RunID, runevent.Phase, runevent.Result, inDataJSON, runevent.DataVersion)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update runEvent")
	}

	return res, nil
}

func (d *DB) insertRawRunEventPostgres(tx *sql.Tx, runevent *types.RunEvent) error {
	inDataJSON, err := json.Marshal(runevent.Data)
	if err != nil {
		return errors.Wrap(err, "failed to marshal runevent.Data")
	}
	q := runEventInsertRawPostgres(runevent.ID, runevent.Revision, runevent.CreationTime, runevent.UpdateTime, runevent.Sequence, runevent.RunEventType, runevent.RunID, runevent.Phase, runevent.Result, inDataJSON, runevent.DataVersion)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert runEvent")
	}

	return nil
}
var (
	executorInsertPostgres = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inExecutorID string, inListenURL string, inArchs []byte, inLabels []byte, inAllowPrivilegedContainers bool, inActiveTasksLimit int, inActiveTasks int, inDynamic bool, inExecutorGroup string, inSiblingsExecutors []byte) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("executor").Cols("id", "revision", "creation_time", "update_time", "executor_id", "listen_url", "archs", "labels", "allow_privileged_containers", "active_tasks_limit", "active_tasks", "dynamic", "executor_group", "siblings_executors").Values(inID, inRevision, inCreationTime, inUpdateTime, inExecutorID, inListenURL, inArchs, inLabels, inAllowPrivilegedContainers, inActiveTasksLimit, inActiveTasks, inDynamic, inExecutorGroup, inSiblingsExecutors)
	}
	executorUpdatePostgres = func(curRevision uint64, inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inExecutorID string, inListenURL string, inArchs []byte, inLabels []byte, inAllowPrivilegedContainers bool, inActiveTasksLimit int, inActiveTasks int, inDynamic bool, inExecutorGroup string, inSiblingsExecutors []byte) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("executor").Set(ub.Assign("id", inID), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("executor_id", inExecutorID), ub.Assign("listen_url", inListenURL), ub.Assign("archs", inArchs), ub.Assign("labels", inLabels), ub.Assign("allow_privileged_containers", inAllowPrivilegedContainers), ub.Assign("active_tasks_limit", inActiveTasksLimit), ub.Assign("active_tasks", inActiveTasks), ub.Assign("dynamic", inDynamic), ub.Assign("executor_group", inExecutorGroup), ub.Assign("siblings_executors", inSiblingsExecutors)).Where(ub.E("id", inID), ub.E("revision", curRevision))
	}

	executorInsertRawPostgres = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inExecutorID string, inListenURL string, inArchs []byte, inLabels []byte, inAllowPrivilegedContainers bool, inActiveTasksLimit int, inActiveTasks int, inDynamic bool, inExecutorGroup string, inSiblingsExecutors []byte) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("executor").Cols("id", "revision", "creation_time", "update_time", "executor_id", "listen_url", "archs", "labels", "allow_privileged_containers", "active_tasks_limit", "active_tasks", "dynamic", "executor_group", "siblings_executors").SQL("OVERRIDING SYSTEM VALUE").Values(inID, inRevision, inCreationTime, inUpdateTime, inExecutorID, inListenURL, inArchs, inLabels, inAllowPrivilegedContainers, inActiveTasksLimit, inActiveTasks, inDynamic, inExecutorGroup, inSiblingsExecutors)
	}
)

func (d *DB) insertExecutorPostgres(tx *sql.Tx, executor *types.Executor) error {
	inArchsJSON, err := json.Marshal(executor.Archs)
	if err != nil {
		return errors.Wrap(err, "failed to marshal executor.Archs")
	}
	inLabelsJSON, err := json.Marshal(executor.Labels)
	if err != nil {
		return errors.Wrap(err, "failed to marshal executor.Labels")
	}
	inSiblingsExecutorsJSON, err := json.Marshal(executor.SiblingsExecutors)
	if err != nil {
		return errors.Wrap(err, "failed to marshal executor.SiblingsExecutors")
	}
	q := executorInsertPostgres(executor.ID, executor.Revision, executor.CreationTime, executor.UpdateTime, executor.ExecutorID, executor.ListenURL, inArchsJSON, inLabelsJSON, executor.AllowPrivilegedContainers, executor.ActiveTasksLimit, executor.ActiveTasks, executor.Dynamic, executor.ExecutorGroup, inSiblingsExecutorsJSON)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert executor")
	}

	return nil
}

func (d *DB) updateExecutorPostgres(tx *sql.Tx, curRevision uint64, executor *types.Executor) (stdsql.Result, error) {
	inArchsJSON, err := json.Marshal(executor.Archs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal executor.Archs")
	}
	inLabelsJSON, err := json.Marshal(executor.Labels)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal executor.Labels")
	}
	inSiblingsExecutorsJSON, err := json.Marshal(executor.SiblingsExecutors)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal executor.SiblingsExecutors")
	}
	q := executorUpdatePostgres(curRevision, executor.ID, executor.Revision, executor.CreationTime, executor.UpdateTime, executor.ExecutorID, executor.ListenURL, inArchsJSON, inLabelsJSON, executor.AllowPrivilegedContainers, executor.ActiveTasksLimit, executor.ActiveTasks, executor.Dynamic, executor.ExecutorGroup, inSiblingsExecutorsJSON)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update executor")
	}

	return res, nil
}

func (d *DB) insertRawExecutorPostgres(tx *sql.Tx, executor *types.Executor) error {
	inArchsJSON, err := json.Marshal(executor.Archs)
	if err != nil {
		return errors.Wrap(err, "failed to marshal executor.Archs")
	}
	inLabelsJSON, err := json.Marshal(executor.Labels)
	if err != nil {
		return errors.Wrap(err, "failed to marshal executor.Labels")
	}
	inSiblingsExecutorsJSON, err := json.Marshal(executor.SiblingsExecutors)
	if err != nil {
		return errors.Wrap(err, "failed to marshal executor.SiblingsExecutors")
	}
	q := executorInsertRawPostgres(executor.ID, executor.Revision, executor.CreationTime, executor.UpdateTime, executor.ExecutorID, executor.ListenURL, inArchsJSON, inLabelsJSON, executor.AllowPrivilegedContainers, executor.ActiveTasksLimit, executor.ActiveTasks, executor.Dynamic, executor.ExecutorGroup, inSiblingsExecutorsJSON)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert executor")
	}

	return nil
}
var (
	executorTaskInsertPostgres = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inExecutorID string, inRunID string, inRunTaskID string, inStop bool, inPhase types.ExecutorTaskPhase, inTimedout bool, inFailError string, inStartTime *time.Time, inEndTime *time.Time, inSetupStep []byte, inSteps []byte) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("executortask").Cols("id", "revision", "creation_time", "update_time", "executor_id", "run_id", "run_task_id", "stop", "phase", "timedout", "fail_error", "start_time", "end_time", "setup_step", "steps").Values(inID, inRevision, inCreationTime, inUpdateTime, inExecutorID, inRunID, inRunTaskID, inStop, inPhase, inTimedout, inFailError, inStartTime, inEndTime, inSetupStep, inSteps)
	}
	executorTaskUpdatePostgres = func(curRevision uint64, inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inExecutorID string, inRunID string, inRunTaskID string, inStop bool, inPhase types.ExecutorTaskPhase, inTimedout bool, inFailError string, inStartTime *time.Time, inEndTime *time.Time, inSetupStep []byte, inSteps []byte) *sq.UpdateBuilder {
		ub:= sq.NewUpdateBuilder()
		return ub.Update("executortask").Set(ub.Assign("id", inID), ub.Assign("revision", inRevision), ub.Assign("creation_time", inCreationTime), ub.Assign("update_time", inUpdateTime), ub.Assign("executor_id", inExecutorID), ub.Assign("run_id", inRunID), ub.Assign("run_task_id", inRunTaskID), ub.Assign("stop", inStop), ub.Assign("phase", inPhase), ub.Assign("timedout", inTimedout), ub.Assign("fail_error", inFailError), ub.Assign("start_time", inStartTime), ub.Assign("end_time", inEndTime), ub.Assign("setup_step", inSetupStep), ub.Assign("steps", inSteps)).Where(ub.E("id", inID), ub.E("revision", curRevision))
	}

	executorTaskInsertRawPostgres = func(inID string, inRevision uint64, inCreationTime time.Time, inUpdateTime time.Time, inExecutorID string, inRunID string, inRunTaskID string, inStop bool, inPhase types.ExecutorTaskPhase, inTimedout bool, inFailError string, inStartTime *time.Time, inEndTime *time.Time, inSetupStep []byte, inSteps []byte) *sq.InsertBuilder {
		ib:= sq.NewInsertBuilder()
		return ib.InsertInto("executortask").Cols("id", "revision", "creation_time", "update_time", "executor_id", "run_id", "run_task_id", "stop", "phase", "timedout", "fail_error", "start_time", "end_time", "setup_step", "steps").SQL("OVERRIDING SYSTEM VALUE").Values(inID, inRevision, inCreationTime, inUpdateTime, inExecutorID, inRunID, inRunTaskID, inStop, inPhase, inTimedout, inFailError, inStartTime, inEndTime, inSetupStep, inSteps)
	}
)

func (d *DB) insertExecutorTaskPostgres(tx *sql.Tx, executortask *types.ExecutorTask) error {
	inSetupStepJSON, err := json.Marshal(executortask.SetupStep)
	if err != nil {
		return errors.Wrap(err, "failed to marshal executortask.SetupStep")
	}
	inStepsJSON, err := json.Marshal(executortask.Steps)
	if err != nil {
		return errors.Wrap(err, "failed to marshal executortask.Steps")
	}
	q := executorTaskInsertPostgres(executortask.ID, executortask.Revision, executortask.CreationTime, executortask.UpdateTime, executortask.ExecutorID, executortask.RunID, executortask.RunTaskID, executortask.Stop, executortask.Phase, executortask.Timedout, executortask.FailError, executortask.StartTime, executortask.EndTime, inSetupStepJSON, inStepsJSON)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert executorTask")
	}

	return nil
}

func (d *DB) updateExecutorTaskPostgres(tx *sql.Tx, curRevision uint64, executortask *types.ExecutorTask) (stdsql.Result, error) {
	inSetupStepJSON, err := json.Marshal(executortask.SetupStep)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal executortask.SetupStep")
	}
	inStepsJSON, err := json.Marshal(executortask.Steps)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal executortask.Steps")
	}
	q := executorTaskUpdatePostgres(curRevision, executortask.ID, executortask.Revision, executortask.CreationTime, executortask.UpdateTime, executortask.ExecutorID, executortask.RunID, executortask.RunTaskID, executortask.Stop, executortask.Phase, executortask.Timedout, executortask.FailError, executortask.StartTime, executortask.EndTime, inSetupStepJSON, inStepsJSON)

	res, err := d.exec(tx, q)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update executorTask")
	}

	return res, nil
}

func (d *DB) insertRawExecutorTaskPostgres(tx *sql.Tx, executortask *types.ExecutorTask) error {
	inSetupStepJSON, err := json.Marshal(executortask.SetupStep)
	if err != nil {
		return errors.Wrap(err, "failed to marshal executortask.SetupStep")
	}
	inStepsJSON, err := json.Marshal(executortask.Steps)
	if err != nil {
		return errors.Wrap(err, "failed to marshal executortask.Steps")
	}
	q := executorTaskInsertRawPostgres(executortask.ID, executortask.Revision, executortask.CreationTime, executortask.UpdateTime, executortask.ExecutorID, executortask.RunID, executortask.RunTaskID, executortask.Stop, executortask.Phase, executortask.Timedout, executortask.FailError, executortask.StartTime, executortask.EndTime, inSetupStepJSON, inStepsJSON)

	if _, err := d.exec(tx, q); err != nil {
		return errors.Wrap(err, "failed to insert executorTask")
	}

	return nil
}
