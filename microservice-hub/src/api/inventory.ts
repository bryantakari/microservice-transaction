import { inventoryApi } from "./service";
import type { InsertItemsRequest, ItemsDTO, ListItemsQuery, DataResponse } from "@/types/inventory";

export function insertItems(request: InsertItemsRequest) {
  return inventoryApi.post<DataResponse<ItemsDTO>>("/item", request);
}

export function listItems(request: ListItemsQuery) {
  return inventoryApi.get<DataResponse<ItemsDTO[]>>("/item/list",{ params: request });
}