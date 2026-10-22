// Define the structure for the request payload to insert items
export interface InsertItemsRequest {
  items: ItemsDTO[];
}

export interface ItemsDTO {
  item_name: string;
  item_description: string;
  price: number;   // int64 in Go → number in TS
  item_code: string;
  quantity: number;
}
export interface DataResponse<T>{
    data: T;
}

export interface InsertItemsRequest {
  items: ItemsDTO[];
}
 
export interface ListItemsQuery {
  page: number;
  limit: number;
  sort_by: string;
  order_by: "asc" | "desc"; // restrict to asc | desc
}