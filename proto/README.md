<!--
MIT License

Source: https://gist.github.com/LeKovr/d6ee7d31c65a4b7e90d8d94295e4d535
Copyright (c) 2021 Aleksei Kovrizhkin (LeKovr)

Original version: https://github.com/pseudomuto/protoc-gen-doc/blob/v1.4.1/resources/markdown.tmpl
Copyright (c) 2017 David Muto (pseudomuto)

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
-->

<!-- use first file package name -->
# api.showonce.v1 API Documentation

<a name="top"></a>

## Table of Contents

- [proto/service.proto](#proto/service.proto)
  - Services
      - [PrivateService](#api.showonce.v1.PrivateService)
      - [PublicService](#api.showonce.v1.PublicService)
  
  - Messages
      - [ItemData](#api.showonce.v1.ItemData)
      - [ItemId](#api.showonce.v1.ItemId)
      - [ItemList](#api.showonce.v1.ItemList)
      - [ItemMeta](#api.showonce.v1.ItemMeta)
      - [ItemMetaWithId](#api.showonce.v1.ItemMetaWithId)
      - [NewItemRequest](#api.showonce.v1.NewItemRequest)
      - [Stats](#api.showonce.v1.Stats)
      - [StatsResponse](#api.showonce.v1.StatsResponse)
  
  - Enums
      - [ItemStatus](#api.showonce.v1.ItemStatus)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto/service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/service.proto




<a name="api.showonce.v1.PrivateService"></a>

### PrivateService

Private

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| NewItem | [NewItemRequest](#api.showonce.v1.NewItemRequest) | [ItemId](#api.showonce.v1.ItemId) | создать секрет |
| GetItems | [.google.protobuf.Empty](#google.protobuf.Empty) | [ItemList](#api.showonce.v1.ItemList) | вернуть список своих секретов |
| GetStats | [.google.protobuf.Empty](#google.protobuf.Empty) | [StatsResponse](#api.showonce.v1.StatsResponse) | общая статистика по количеству секретов |


<a name="api.showonce.v1.PublicService"></a>

### PublicService



| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetMetadata | [ItemId](#api.showonce.v1.ItemId) | [ItemMeta](#api.showonce.v1.ItemMeta) | вернуть метаданные секрета по id |
| GetData | [ItemId](#api.showonce.v1.ItemId) | [ItemData](#api.showonce.v1.ItemData) | вернуть текст секрета по id |

 <!-- end services -->


<a name="api.showonce.v1.ItemData"></a>

### ItemData

Данные секрета


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [string](#string) |  | Данные секрета |



          


<a name="api.showonce.v1.ItemId"></a>

### ItemId

Идентификатор (ULID)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |



          


<a name="api.showonce.v1.ItemList"></a>

### ItemList

Список секретов


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [ItemMetaWithId](#api.showonce.v1.ItemMetaWithId) | repeated | Список секретов |



          


<a name="api.showonce.v1.ItemMeta"></a>

### ItemMeta

Метаданные секрета


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  | описание |
| group | [string](#string) |  | идентификатор для группировки |
| owner | [string](#string) |  | автор |
| status | [ItemStatus](#api.showonce.v1.ItemStatus) |  | статус |
| created_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | момент создания |
| modified_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | первоначально - срок автоудаления, после показа - момент показа |



          


<a name="api.showonce.v1.ItemMetaWithId"></a>

### ItemMetaWithId

Метаданные секрета с идентификатором


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | идентификатор |
| meta | [ItemMeta](#api.showonce.v1.ItemMeta) |  | метаданные |



          


<a name="api.showonce.v1.NewItemRequest"></a>

### NewItemRequest

Аргументы запроса на создание


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  | описание |
| group | [string](#string) |  | идентификатор для группировки |
| expire | [string](#string) |  | срок актуальности |
| expire_unit | [string](#string) |  | единица срока актуальности ("d" - день, остальные варианты - как в go: "ns", "us" (or "µs"), "ms", "s", "m", "h") |
| data | [string](#string) |  | текст секрета, удаляется после первого показа |



          


<a name="api.showonce.v1.Stats"></a>

### Stats

Статистика по секретам


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int32](#int32) |  | Всего в хранилище |
| wait | [int32](#int32) |  | Готово к прочтению |
| read | [int32](#int32) |  | Прочитано |
| expired | [int32](#int32) |  | Истек срок актуальности |



          


<a name="api.showonce.v1.StatsResponse"></a>

### StatsResponse

Ответ на запрос статистика


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| my | [Stats](#api.showonce.v1.Stats) |  | Данные по текущему пользователю |
| other | [Stats](#api.showonce.v1.Stats) |  | Данные по остальным пользователям |



          

 <!-- end messages -->


<a name="api.showonce.v1.ItemStatus"></a>

### ItemStatus

Статус секрета

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | A Standard tournament |
| WAIT | 1 | Готово к прочтению |
| READ | 2 | Прочитано |
| EXPIRED | 3 | Истек срок актуальности |
| CLEARED | 4 | Удалено |


 <!-- end enums -->

 <!-- end HasExtensions -->



## Scalar Value Types

| .proto Type | Notes | Go  | C++  | Java |
| ----------- | ----- | --- | ---- | ---- |
| <a name="double" /> double |  | float64 | double | double |
| <a name="float" /> float |  | float32 | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int32 | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | int64 | long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | uint32 | int |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | uint64 | long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int32 | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | int64 | long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | uint32 | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | uint64 | long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int32 | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | int64 | long |
| <a name="bool" /> bool |  | bool | bool | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | string | String |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | []byte | string | ByteString |

