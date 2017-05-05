gorm-explain
=========================

extension library of gorm. explain log output.

# Usage

```go
db, _ := gorm.Open("mysql", buildDataSourceName(opts))
defer db.Close()
db.Callback().Query().Register("explain", gorm_explain.Callback) // Add Callback 
```

# Example
[example code](example/main.go)

## Output Sample

```
$ go run main.go
(/Users/kyokomi/workspace/go/src/github.com/kyokomi/gorm-explain/example/main.go:53) 
[2017-05-05 18:15:44]  [1.02ms]  INSERT INTO `products` (`created_at`,`updated_at`,`deleted_at`,`code`,`price`) VALUES ('2017-05-05 18:15:44','2017-05-05 18:15:44',NULL,'L1212','1000')

(/Users/kyokomi/workspace/go/src/github.com/kyokomi/gorm-explain/example/main.go:57) 
[2017-05-05 18:15:44]  [1.25ms]  SELECT * FROM `products`  WHERE `products`.`deleted_at` IS NULL AND ((`products`.`id` = '1')) ORDER BY `products`.`id` ASC LIMIT 1
+-------+----------------+----------+---------------+---------+------------------+--------+------------+--------+---------+-------------+--------------------------------------------------------+
|    id |    select_type |    table |    partitions |    type |    possible_keys |    key |    key_len |    ref |    rows |    filtered |                                                  Extra |
+=======+================+==========+===============+=========+==================+========+============+========+=========+=============+========================================================+
|     1 |         SIMPLE |          |               |         |                  |        |            |        |         |             |    Impossible WHERE noticed after reading const tables |
+-------+----------------+----------+---------------+---------+------------------+--------+------------+--------+---------+-------------+--------------------------------------------------------+


(/Users/kyokomi/workspace/go/src/github.com/kyokomi/gorm-explain/example/main.go:58) 
[2017-05-05 18:15:44]  [1.67ms]  SELECT * FROM `products`  WHERE `products`.`deleted_at` IS NULL AND ((code = 'L1212')) ORDER BY `products`.`id` ASC LIMIT 1
+-------+----------------+-------------+---------------+---------+----------------------------+----------------------------+------------+----------+---------+-------------+---------------------------------------+
|    id |    select_type |       table |    partitions |    type |              possible_keys |                        key |    key_len |      ref |    rows |    filtered |                                 Extra |
+=======+================+=============+===============+=========+============================+============================+============+==========+=========+=============+=======================================+
|     1 |         SIMPLE |    products |               |     ref |    idx_products_deleted_at |    idx_products_deleted_at |          5 |    const |       1 |         100 |    Using index condition; Using where |
+-------+----------------+-------------+---------------+---------+----------------------------+----------------------------+------------+----------+---------+-------------+---------------------------------------+


(/Users/kyokomi/workspace/go/src/github.com/kyokomi/gorm-explain/example/main.go:61) 
[2017-05-05 18:15:44]  [0.90ms]  UPDATE `products` SET `price` = '2000', `updated_at` = '2017-05-05 18:15:44'  WHERE `products`.`deleted_at` IS NULL AND `products`.`id` = '6'

(/Users/kyokomi/workspace/go/src/github.com/kyokomi/gorm-explain/example/main.go:64) 
[2017-05-05 18:15:44]  [1.27ms]  UPDATE `products` SET `deleted_at`='2017-05-05 18:15:44'  WHERE `products`.`deleted_at` IS NULL AND `products`.`id` = '6'
```

# Support Driver
- mysql

# Requirement
- [jinzhu/gorm](https://github.com/jinzhu/gorm)
- [bndr/gotabulate](github.com/bndr/gotabulate)

# License

[MIT](LICENSE)
