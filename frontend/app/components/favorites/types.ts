import { ItemExtended } from "@/proto/app/protos/reqHandlerTypes"
import { PrefForm } from "@/types/filterTypes"

export interface ItemWithFilters {
  item: ItemExtended
  filters: PrefForm
  market: number
  categoryURL: string
}

export interface FavRecord {
  [key: string]: ItemWithFilters
}