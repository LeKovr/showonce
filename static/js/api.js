var AppAPI = (() => {
  var __defProp = Object.defineProperty;
  var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
  var __getOwnPropNames = Object.getOwnPropertyNames;
  var __hasOwnProp = Object.prototype.hasOwnProperty;
  var __export = (target, all) => {
    for (var name in all)
      __defProp(target, name, { get: all[name], enumerable: true });
  };
  var __copyProps = (to, from, except, desc) => {
    if (from && typeof from === "object" || typeof from === "function") {
      for (let key of __getOwnPropNames(from))
        if (!__hasOwnProp.call(to, key) && key !== except)
          __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
    }
    return to;
  };
  var __toCommonJS = (mod) => __copyProps(__defProp({}, "__esModule", { value: true }), mod);

  // zgen/ts/proto/service.pb.ts
  var service_pb_exports = {};
  __export(service_pb_exports, {
    ItemStatus: () => ItemStatus,
    PrivateService: () => PrivateService,
    PublicService: () => PublicService
  });

  // zgen/ts/fetch.pb.ts
  var b64 = new Array(64);
  var s64 = new Array(123);
  for (let i = 0; i < 64; )
    s64[b64[i] = i < 26 ? i + 65 : i < 52 ? i + 71 : i < 62 ? i - 4 : i - 59 | 43] = i++;
  function b64Encode(buffer, start, end) {
    let parts = null;
    const chunk = [];
    let i = 0, j = 0, t;
    while (start < end) {
      const b = buffer[start++];
      switch (j) {
        case 0:
          chunk[i++] = b64[b >> 2];
          t = (b & 3) << 4;
          j = 1;
          break;
        case 1:
          chunk[i++] = b64[t | b >> 4];
          t = (b & 15) << 2;
          j = 2;
          break;
        case 2:
          chunk[i++] = b64[t | b >> 6];
          chunk[i++] = b64[b & 63];
          j = 0;
          break;
      }
      if (i > 8191) {
        (parts || (parts = [])).push(String.fromCharCode.apply(String, chunk));
        i = 0;
      }
    }
    if (j) {
      chunk[i++] = b64[t];
      chunk[i++] = 61;
      if (j === 1)
        chunk[i++] = 61;
    }
    if (parts) {
      if (i)
        parts.push(String.fromCharCode.apply(String, chunk.slice(0, i)));
      return parts.join("");
    }
    return String.fromCharCode.apply(String, chunk.slice(0, i));
  }
  function replacer(key, value) {
    if (value && value.constructor === Uint8Array) {
      return b64Encode(value, 0, value.length);
    }
    return value;
  }
  function fetchReq(path, init) {
    const { pathPrefix, ...req } = init || {};
    const url = pathPrefix ? `${pathPrefix}${path}` : path;
    return fetch(url, req).then((r) => r.json().then((body) => {
      if (!r.ok) {
        throw body;
      }
      return body;
    }));
  }
  function isPlainObject(value) {
    const isObject = Object.prototype.toString.call(value).slice(8, -1) === "Object";
    const isObjLike = value !== null && isObject;
    if (!isObjLike || !isObject) {
      return false;
    }
    const proto = Object.getPrototypeOf(value);
    const hasObjectConstructor = typeof proto === "object" && proto.constructor === Object.prototype.constructor;
    return hasObjectConstructor;
  }
  function isPrimitive(value) {
    return ["string", "number", "boolean"].some((t) => typeof value === t);
  }
  function isZeroValuePrimitive(value) {
    return value === false || value === 0 || value === "";
  }
  function flattenRequestPayload(requestPayload, path = "") {
    return Object.keys(requestPayload).reduce(
      (acc, key) => {
        const value = requestPayload[key];
        const newPath = path ? [path, key].join(".") : key;
        const isNonEmptyPrimitiveArray = Array.isArray(value) && value.every((v) => isPrimitive(v)) && value.length > 0;
        const isNonZeroValuePrimitive = isPrimitive(value) && !isZeroValuePrimitive(value);
        let objectToMerge = {};
        if (isPlainObject(value)) {
          objectToMerge = flattenRequestPayload(value, newPath);
        } else if (isNonZeroValuePrimitive || isNonEmptyPrimitiveArray) {
          objectToMerge = { [newPath]: value };
        }
        return { ...acc, ...objectToMerge };
      },
      {}
    );
  }
  function renderURLSearchParams(requestPayload, urlPathParams = []) {
    const flattenedRequestPayload = flattenRequestPayload(requestPayload);
    const urlSearchParams = Object.keys(flattenedRequestPayload).reduce(
      (acc, key) => {
        const value = flattenedRequestPayload[key];
        if (urlPathParams.find((f) => f === key)) {
          return acc;
        }
        return Array.isArray(value) ? [...acc, ...value.map((m) => [key, m.toString()])] : acc = [...acc, [key, value.toString()]];
      },
      []
    );
    return new URLSearchParams(urlSearchParams).toString();
  }

  // zgen/ts/proto/service.pb.ts
  var ItemStatus = /* @__PURE__ */ ((ItemStatus2) => {
    ItemStatus2["UNKNOWN"] = "UNKNOWN";
    ItemStatus2["WAIT"] = "WAIT";
    ItemStatus2["READ"] = "READ";
    ItemStatus2["EXPIRED"] = "EXPIRED";
    ItemStatus2["CLEARED"] = "CLEARED";
    return ItemStatus2;
  })(ItemStatus || {});
  var PublicService = class {
    static GetMetadata(req, initReq) {
      return fetchReq(`/api/item?${renderURLSearchParams(req, [])}`, { ...initReq, method: "GET" });
    }
    static GetData(req, initReq) {
      return fetchReq(`/api/item/${req["id"]}`, { ...initReq, method: "POST" });
    }
  };
  var PrivateService = class {
    static NewItem(req, initReq) {
      return fetchReq(`/my/api/new`, { ...initReq, method: "POST", body: JSON.stringify(req, replacer) });
    }
    static GetItems(req, initReq) {
      return fetchReq(`/my/api/items?${renderURLSearchParams(req, [])}`, { ...initReq, method: "GET" });
    }
    static GetStats(req, initReq) {
      return fetchReq(`/my/api/stat?${renderURLSearchParams(req, [])}`, { ...initReq, method: "GET" });
    }
  };
  return __toCommonJS(service_pb_exports);
})();
