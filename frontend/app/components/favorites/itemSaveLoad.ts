import { getFav, saveFav } from "@/components/favorites/favSaveLoad";
import { FavRecord, ItemWithFilters } from "@/components/favorites/types";

export function saveToFav(data: ItemWithFilters, itemUrl: string): void {
  const favorites: FavRecord = getFav();
  favorites[itemUrl] = data;
  saveFav(favorites);
}

export function removeFromFav(itemUrl: string): void {
  const favorites: FavRecord = getFav();
  delete favorites[itemUrl];
  saveFav(favorites);
}

export function getFromFav(itemUrl: string): ItemWithFilters | undefined {
  const favorites: FavRecord = getFav();
  return favorites[itemUrl];
}
