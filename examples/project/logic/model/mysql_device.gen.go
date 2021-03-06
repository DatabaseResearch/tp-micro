// Code generated by 'micro gen' command.
// DO NOT EDIT!

package model

import (
	"database/sql"
	"unsafe"

	"github.com/henrylee2cn/goutil/coarsetime"
	tp "github.com/henrylee2cn/teleport"
	"github.com/xiaoenai/tp-micro/model/mysql"
	"github.com/xiaoenai/tp-micro/model/sqlx"

	"github.com/xiaoenai/tp-micro/examples/project/args"
)

// Device comment...
type Device args.Device

// ToDevice converts to *Device type.
func ToDevice(_d *args.Device) *Device {
	return (*Device)(unsafe.Pointer(_d))
}

// ToArgsDevice converts to *args.Device type.
func ToArgsDevice(_d *Device) *args.Device {
	return (*args.Device)(unsafe.Pointer(_d))
}

// ToDeviceSlice converts to []*Device type.
func ToDeviceSlice(a []*args.Device) []*Device {
	return *(*[]*Device)(unsafe.Pointer(&a))
}

// ToArgsDeviceSlice converts to []*args.Device type.
func ToArgsDeviceSlice(a []*Device) []*args.Device {
	return *(*[]*args.Device)(unsafe.Pointer(&a))
}

// TableName implements 'github.com/xiaoenai/tp-micro/model'.Cacheable
func (*Device) TableName() string {
	return "device"
}

func (_d *Device) isZeroPrimaryKey() bool {
	var _uuid string
	if _d.UUID != _uuid {
		return false
	}
	return true
}

var deviceDB, _ = mysqlHandler.RegCacheableDB(new(Device), cacheExpire, args.DeviceSql)

// GetDeviceDB returns the Device DB handler.
func GetDeviceDB() *mysql.CacheableDB {
	return deviceDB
}

// InsertDevice insert a Device data into database.
// NOTE:
//  Primary key: 'uuid';
//  Without cache layer.
func InsertDevice(_d *Device, tx ...*sqlx.Tx) error {
	_d.UpdatedAt = coarsetime.FloorTimeNow().Unix()
	if _d.CreatedAt == 0 {
		_d.CreatedAt = _d.UpdatedAt
	}
	return deviceDB.Callback(func(tx sqlx.DbOrTx) error {
		var (
			query            string
			isZeroPrimaryKey = _d.isZeroPrimaryKey()
		)
		if isZeroPrimaryKey {
			query = "INSERT INTO `device` (`updated_at`,`created_at`)VALUES(:updated_at,:created_at);"
		} else {
			query = "INSERT INTO `device` (`uuid`,`updated_at`,`created_at`)VALUES(:uuid,:updated_at,:created_at);"
		}
		_, err := tx.NamedExec(query, _d)
		return err
	}, tx...)
}

// UpsertDevice insert or update the Device data by primary key.
// NOTE:
//  Primary key: 'uuid';
//  With cache layer;
//  Insert data if the primary key is specified;
//  Update data based on _updateFields if no primary key is specified;
//  _updateFields' members must be db field style (snake format);
//  Automatic update 'updated_at' field;
//  Don't update the primary keys, 'created_at' key and 'deleted_ts' key;
//  Update all fields except the primary keys, 'created_at' key and 'deleted_ts' key, if _updateFields is empty.
func UpsertDevice(_d *Device, _updateFields []string, tx ...*sqlx.Tx) error {
	if _d.UpdatedAt == 0 {
		_d.UpdatedAt = coarsetime.FloorTimeNow().Unix()
	}
	if _d.CreatedAt == 0 {
		_d.CreatedAt = _d.UpdatedAt
	}
	err := deviceDB.Callback(func(tx sqlx.DbOrTx) error {
		var (
			query            string
			isZeroPrimaryKey = _d.isZeroPrimaryKey()
		)
		if isZeroPrimaryKey {
			query = "INSERT INTO `device` (`updated_at`,`created_at`)VALUES(:updated_at,:created_at)"
		} else {
			query = "INSERT INTO `device` (`uuid`,`updated_at`,`created_at`)VALUES(:uuid,:updated_at,:created_at)"
		}
		query += " ON DUPLICATE KEY UPDATE "
		if len(_updateFields) == 0 {
			query += "`updated_at`=VALUES(`updated_at`);"
		} else {
			for _, s := range _updateFields {
				if s == "updated_at" || s == "created_at" || s == "deleted_ts" || s == "uuid" {
					continue
				}
				query += "`" + s + "`=VALUES(`" + s + "`),"
			}
			if query[len(query)-1] != ',' {
				return nil
			}
			query += "`updated_at`=VALUES(`updated_at`),`deleted_ts`=0;"
		}
		_, err := tx.NamedExec(query, _d)
		return err
	}, tx...)
	if err != nil {
		return err
	}
	err = deviceDB.DeleteCache(_d)
	if err != nil {
		tp.Errorf("%s", err.Error())
	}
	return nil
}

// UpdateDeviceByPrimary update the Device data in database by primary key.
// NOTE:
//  Primary key: 'uuid';
//  With cache layer;
//  _updateFields' members must be db field style (snake format);
//  Automatic update 'updated_at' field;
//  Don't update the primary keys, 'created_at' key and 'deleted_ts' key;
//  Update all fields except the primary keys, 'created_at' key and 'deleted_ts' key, if _updateFields is empty.
func UpdateDeviceByPrimary(_d *Device, _updateFields []string, tx ...*sqlx.Tx) error {
	_d.UpdatedAt = coarsetime.FloorTimeNow().Unix()
	err := deviceDB.Callback(func(tx sqlx.DbOrTx) error {
		query := "UPDATE `device` SET "
		if len(_updateFields) == 0 {
			query += "`updated_at`=:updated_at WHERE `uuid`=:uuid AND `deleted_ts`=0 LIMIT 1;"
		} else {
			for _, s := range _updateFields {
				if s == "updated_at" || s == "created_at" || s == "deleted_ts" || s == "uuid" {
					continue
				}
				query += "`" + s + "`=:" + s + ","
			}
			if query[len(query)-1] != ',' {
				return nil
			}
			query += "`updated_at`=:updated_at WHERE `uuid`=:uuid AND `deleted_ts`=0 LIMIT 1;"
		}
		_, err := tx.NamedExec(query, _d)
		return err
	}, tx...)
	if err != nil {
		return err
	}
	err = deviceDB.DeleteCache(_d)
	if err != nil {
		tp.Errorf("%s", err.Error())
	}
	return nil
}

// DeleteDeviceByPrimary delete a Device data in database by primary key.
// NOTE:
//  Primary key: 'uuid';
//  With cache layer.
func DeleteDeviceByPrimary(_uuid string, deleteHard bool, tx ...*sqlx.Tx) error {
	var err error
	if deleteHard {
		// Immediately delete from the hard disk.
		err = deviceDB.Callback(func(tx sqlx.DbOrTx) error {
			_, err := tx.Exec("DELETE FROM `device` WHERE `uuid`=? AND `deleted_ts`=0;", _uuid)
			return err
		}, tx...)

	} else {
		// Delay delete from the hard disk.
		ts := coarsetime.FloorTimeNow().Unix()
		err = deviceDB.Callback(func(tx sqlx.DbOrTx) error {
			_, err := tx.Exec("UPDATE `device` SET `updated_at`=?, `deleted_ts`=? WHERE `uuid`=? AND `deleted_ts`=0;", ts, ts, _uuid)
			return err
		}, tx...)
	}

	if err != nil {
		return err
	}
	err = deviceDB.DeleteCache(&Device{
		UUID: _uuid,
	})
	if err != nil {
		tp.Errorf("%s", err.Error())
	}
	return nil
}

// GetDeviceByPrimary query a Device data from database by primary key.
// NOTE:
//  Primary key: 'uuid';
//  With cache layer;
//  If @return bool=false error=nil, means the data is not exist.
func GetDeviceByPrimary(_uuid string) (*Device, bool, error) {
	var _d = &Device{
		UUID: _uuid,
	}
	err := deviceDB.CacheGet(_d)
	switch err {
	case nil:
		if _d.CreatedAt == 0 {
			return nil, false, nil
		}
		return _d, true, nil
	case sql.ErrNoRows:
		return nil, false, nil
	default:
		return nil, false, err
	}
}

// GetDeviceByWhere query a Device data from database by WHERE condition.
// NOTE:
//  Without cache layer;
//  If @return bool=false error=nil, means the data is not exist.
func GetDeviceByWhere(whereCond string, arg ...interface{}) (*Device, bool, error) {
	var _d = new(Device)
	err := deviceDB.Get(_d, "SELECT `uuid`,`updated_at`,`created_at` FROM `device` WHERE "+insertZeroDeletedTsField(whereCond)+" LIMIT 1;", arg...)
	switch err {
	case nil:
		return _d, true, nil
	case sql.ErrNoRows:
		return nil, false, nil
	default:
		return nil, false, err
	}
}

// SelectDeviceByWhere query some Device data from database by WHERE condition.
// NOTE:
//  Without cache layer.
func SelectDeviceByWhere(whereCond string, arg ...interface{}) ([]*Device, error) {
	var objs = new([]*Device)
	err := deviceDB.Select(objs, "SELECT `uuid`,`updated_at`,`created_at` FROM `device` WHERE "+insertZeroDeletedTsField(whereCond), arg...)
	return *objs, err
}

// CountDeviceByWhere count Device data number from database by WHERE condition.
// NOTE:
//  Without cache layer.
func CountDeviceByWhere(whereCond string, arg ...interface{}) (int64, error) {
	var count int64
	err := deviceDB.Get(&count, "SELECT count(1) FROM `device` WHERE "+insertZeroDeletedTsField(whereCond), arg...)
	return count, err
}
