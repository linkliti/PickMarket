import { ItemExtended } from "@/proto/app/protos/reqHandlerTypes"
import { PrefForm } from "@/types/filterTypes"

export interface ItemWithFilters {
  item: ItemExtended
  filters: PrefForm
}

export interface FavRecord {
  [key: string]: ItemWithFilters
}