package note

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	leveldbUtil "github.com/syndtr/goleveldb/leveldb/util"
)

//11.1 LevelDB的基本使用
func LeveldbBasic() {
	// 打开数据库
	db, err := leveldb.OpenFile("leveldb", nil)
	if err != nil {
		panic(err)
	}
	// 关闭数据库
	defer db.Close()
	db.Put([]byte("user-1"), []byte("{\"username\":\"1\"}"), nil)
	//db.Delete([]byte("user-1"), nil)

	data1, _ := db.Get([]byte("user-1"), nil)
	fmt.Println("data1=", string(data1))
	//批量写数据
	batch := new(leveldb.Batch)
	batch.Put([]byte("user-2"), []byte("{\"username\":\"2\"}"))
	batch.Put([]byte("user-3"), []byte("{\"username\":\"3\"}"))
	db.Delete([]byte("user-1"), nil)
	batch.Put([]byte("user-1"), []byte("{\"username\":\"11\"}"))
	db.Write(batch, nil)
	data3, _ := db.Get([]byte("user-3"), nil)
	fmt.Println("data=", string(data3))
	data11, _ := db.Get([]byte("user-1"), nil)
	fmt.Println("data1-=", string(data11))

}

//11.1.4 LevelDB遍历
func LeveldbIterate() {
	// 打开数据库
	db, err := leveldb.OpenFile("leveldb", nil)
	if err != nil {
		panic(err)
	}
	// 关闭数据库
	defer db.Close()
	//批量写数据
	batch := new(leveldb.Batch)
	for i := 1; i < 11; i++ {
		batch.Put([]byte(fmt.Sprintf("user-%v", i)), []byte(fmt.Sprintf("{\"name\":\"u%v\"}", i)))
	}
	db.Write(batch, nil)
	iter := db.NewIterator(&leveldbUtil.Range{Start: []byte("user-1"), Limit: []byte("user-8")}, nil)
	for iter.Next() {
		fmt.Printf("%v=%v\n", string(iter.Key()), string(iter.Value()))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	iter = db.NewIterator(leveldbUtil.BytesPrefix([]byte("user-")), nil)
	for iter.Next() {
		fmt.Printf("%v=%v\n", string(iter.Key()), string(iter.Value()))
	}
	iter.Release()
	err = iter.Error()
	if err != nil {
		fmt.Println(err)
	}
}
