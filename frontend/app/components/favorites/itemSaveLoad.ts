import { getFav, saveFav } from "@/components/favorites/favSaveLoad";
import { FavRecord, ItemWithFilters } from "@/components/favorites/types";
import { ItemExtended } from "@/proto/app/protos/reqHandlerTypes";
import { PrefForm } from "@/types/filterTypes";

export function saveToFav({ item, filters }: { item: ItemExtended; filters: PrefForm }): void {
  const key: string | undefined = item.item?.url;
  if (!key) {
    return;
  }
  const itemWithFilters: ItemWithFilters = {
    item,
    filters,
  };
  const favorites: FavRecord = getFav();
  favorites[key] = itemWithFilters;
  saveFav(favorites);
}

export function removeFromFav({ itemUrl }: { itemUrl: string }): void {
  const favorites: FavRecord = getFav();
  delete favorites[itemUrl];
  saveFav(favorites);
}

export function getFromFav({ itemUrl }: { itemUrl: string }): ItemWithFilters | undefined {
  const favorites: FavRecord = getFav();
  return favorites[itemUrl];
}
