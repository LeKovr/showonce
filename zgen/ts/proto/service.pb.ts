/* eslint-disable */
// @ts-nocheck
/*
* This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
*/

import * as fm from "../fetch.pb"
import * as GoogleProtobufEmpty from "../google/protobuf/empty.pb"
import * as GoogleProtobufTimestamp from "../google/protobuf/timestamp.pb"

export enum ItemStatus {
  UNKNOWN = "UNKNOWN",
  WAIT = "WAIT",
  READ = "READ",
  EXPIRED = "EXPIRED",
  CLEARED = "CLEARED",
}

export type ItemId = {
  id?: string
}

export type NewItemRequest = {
  title?: string
  group?: string
  expire?: string
  expireUnit?: string
  data?: string
}

export type ItemData = {
  data?: string
}

export type ItemMeta = {
  title?: string
  group?: string
  owner?: string
  status?: ItemStatus
  createdAt?: GoogleProtobufTimestamp.Timestamp
  modifiedAt?: GoogleProtobufTimestamp.Timestamp
}

export type ItemMetaWithId = {
  id?: string
  meta?: ItemMeta
}

export type ItemList = {
  items?: ItemMetaWithId[]
}

export type Stats = {
  total?: number
  wait?: number
  read?: number
  expired?: number
}

export type StatsResponse = {
  my?: Stats
  other?: Stats
}

export class PublicService {
  static GetMetadata(req: ItemId, initReq?: fm.InitReq): Promise<ItemMeta> {
    return fm.fetchReq<ItemId, ItemMeta>(`/api/item?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetData(req: ItemId, initReq?: fm.InitReq): Promise<ItemData> {
    return fm.fetchReq<ItemId, ItemData>(`/api/item/${req["id"]}`, {...initReq, method: "POST"})
  }
}
export class PrivateService {
  static NewMessage(req: NewItemRequest, initReq?: fm.InitReq): Promise<ItemId> {
    return fm.fetchReq<NewItemRequest, ItemId>(`/my/api/new`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)})
  }
  static GetItems(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<ItemList> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, ItemList>(`/my/api/items?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
  static GetStats(req: GoogleProtobufEmpty.Empty, initReq?: fm.InitReq): Promise<StatsResponse> {
    return fm.fetchReq<GoogleProtobufEmpty.Empty, StatsResponse>(`/my/api/stat?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"})
  }
}