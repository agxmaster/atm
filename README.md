#### 简介
依赖gorm和golang1.8 对gorm做了一点封装看似没比gorm简单,实际也是皮裤套棉裤更麻烦了。 食用方法看test

#### 项目中使用
```
//1. go get this project

//2. gorm.db init
var gormDb *gorm.DB
dbSetetup()

//3. get instance 
db := &atm.Atm[Users]{Db: gormDb}

```

#### condition 支持自定义
1. 可以自己扩展实现Condition.Build方法即可 
2. condition.GormWhere 支持简单的逻辑条件，支持 >= ? , in (?) 等方式和gorm一致
3. condition.Scope 直接使用直接使用gorm 语法即可, 可以实现更复杂的查询逻辑
#### facade
对原有结构体方法进行包装隐藏了泛型结构创建过程，针对在不创建model结构体的情况下，直接操作db，使用方法和atm一致





