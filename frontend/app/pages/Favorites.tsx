import WhiteBlock from "@/components/base/WhiteBlock";
import ItemBlockFav from "@/components/favorites/ItemBlockFav";
import { clearFav, getFav } from "@/components/favorites/favSaveLoad";
import { FavRecord, ItemWithFilters } from "@/components/favorites/types";
import { FilterContextProvider, FilterContextType } from "@/components/items/FilterContext";
import { ItemContextProvider, ItemContextType } from "@/components/items/ItemContext";
import { Button } from "@/components/ui/button";
import { Item } from "@/proto/app/protos/items";
import { ReactElement, useEffect, useState } from "react";

export default function Favorites(): ReactElement {
  useEffect((): void => {
    document.title = "Избранное";
  }, []);

  const [favorites, setFavorites] = useState<FavRecord>({});

  useEffect((): void => {
    const favs: FavRecord = getFav();
    setFavorites(favs);
  }, []);

  if (Object.keys(favorites).length === 0) {
    return <WhiteBlock className="w-full">Нет избранных товаров</WhiteBlock>;
  }

  return (
    <>
      <WhiteBlock className="inline-flex w-full">
        <h2 className="text-2xl font-bold">
          Просмотр избранного ({Object.keys(favorites).length})
        </h2>
        <Button
          className="ml-auto"
          onClick={(): void => {
            clearFav();
            setFavorites({});
          }}
        >
          Очистить
        </Button>
      </WhiteBlock>
      {Object.entries(favorites).map(([key, val]: [string, ItemWithFilters]): ReactElement => {
        const itemBaseData: Item = val.item.item ? val.item.item : ({} as Item);
        const itemData: ItemContextType = {
          item: itemBaseData,
          chars: val.item.chars,
          similar: val.item.similar,
          totalWeight: 0,
          market: val.market,
          maxTotalWeight: 0,
          isFav: true,
          favUrl: key,
        };
        const filterData: FilterContextType = {
          categoryURL: val.categoryURL,
          form: val.filters,
        };
        return (
          <ItemContextProvider
            data={itemData}
            key={key}
          >
            <FilterContextProvider data={filterData}>
              <ItemBlockFav />
            </FilterContextProvider>
          </ItemContextProvider>
        );
      })}
    </>
  );
}
