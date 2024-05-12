import { FavRecord } from "@/components/favorites/types";

export function getFav(): FavRecord {
  const favourites: FavRecord = JSON.parse(localStorage.getItem("fav") || "{}");
  return favourites;
}

export function saveFav(favourites: FavRecord): void {
  localStorage.setItem("fav", JSON.stringify(favourites));
}

export function clearFav(): void {
  localStorage.removeItem("fav");
}