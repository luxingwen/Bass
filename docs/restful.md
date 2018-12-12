# Restful使用文档

### 获取一条记录

```
curl -X GET \
  http://127.0.0.1:8003/api/v1/abc_models/5 \
```

返回

```
{
    "Count": 3,
    "data": [
        {
            "author_id": null,
            "body": null,
            "created_at": null,
            "deleted_at": null,
            "description": null,
            "id": "5",
            "slug": null,
            "title": "334",
            "updated_at": null
        }
    ]
}
```



### 获取所有记录

- [ ] 分页
- [ ] 条件
- [x] 统计
- [ ] 连表查询

```
curl -X GET \
  http://127.0.0.1:8003/api/v1/abc_models \
  -H 'content-type: application/json' \
```

```
{
    "Count": 3,
    "data": [
        {
            "author_id": null,
            "body": null,
            "created_at": null,
            "deleted_at": null,
            "description": null,
            "id": "2",
            "slug": null,
            "title": "4343",
            "updated_at": null
        },
        {
            "author_id": null,
            "body": null,
            "created_at": null,
            "deleted_at": null,
            "description": null,
            "id": "4",
            "slug": null,
            "title": "NULL33",
            "updated_at": null
        }
    ]
}
```



### 增加一条记录

- [ ] 自动创建表
- [ ] 自动创建列

```
curl -X POST \
  http://127.0.0.1:8003/api/v1/abc_models \
  -H 'content-type: application/json' \
  -d '{"title":"test","body":34343}'

返回：
{
    "code": 200,
    "msg": "ok"
}
```



### 删除一条记录

```
curl -X DELETE \
  http://127.0.0.1:8003/api/v1/abc_models/5 \
```



### 修改一条记录

```
curl -X POST \
  http://127.0.0.1:8003/api/v1/abc_models/5 \
  -H 'content-type: application/json' \
  -d '{"title":"test","body":34343}'

返回：
{
    "code": 200,
    "msg": "ok"
}
```

