import WhiteBlock from "@/components/base/WhiteBlock";
import { getFav } from "@/components/favorites/favSaveLoad";
import { FavRecord } from "@/components/favorites/types";
import { FilterContextProvider, FilterContextType } from "@/components/items/FilterContext";
import { ItemContextProvider, ItemContextType } from "@/components/items/ItemContext";
import ItemBlock from "@/components/items/item/ItemBlock";
import useFetchItems from "@/components/items/useFetchItems";
import { Item } from "@/proto/app/protos/items";
import { ItemExtended, UserPref } from "@/proto/app/protos/reqHandlerTypes";
import { FilterStore, useFilterStore } from "@/store/filterStore";
import { PrefForm } from "@/types/filterTypes";
import { LoadingSpinner } from "@/utilities/LoadingSpinner";

import { ReactElement, useEffect } from "react";

export default function ItemsSection(): ReactElement {
  const [activePrefs, market, setActivePrefs, formPrefs, categoryUrl] = useFilterStore(
    (state: FilterStore) => [
      state.activePrefs,
      state.market,
      state.setActivePrefs,
      state.formPrefs,
      state.categoryUrl,
    ],
  );
  const { items, isLoading, error, setItems } = useFetchItems();

  useEffect(() => {
    setItems([]);
    setActivePrefs(null);
  }, [market, setActivePrefs, setItems]);

  if (!activePrefs) {
    return <WhiteBlock className="w-full">Здесь будут отображаться товары</WhiteBlock>;
  }
  if (error) {
    return (
      <WhiteBlock className="w-full">Не удалось получить товары. Попробуйте еще раз</WhiteBlock>
    );
  }
  if (isLoading) {
    return (
      <WhiteBlock className="w-full">
        <div className="flex items-center gap-2">
          <LoadingSpinner /> <p>Загрузка товаров (это может занять некоторое время)</p>
        </div>
      </WhiteBlock>
    );
  }

  const sortedItems: ItemExtended[] = items.sort((a: ItemExtended, b: ItemExtended): 0 | 1 | -1 => {
    if (a.totalWeight > b.totalWeight) {
      return -1;
    }
    if (a.totalWeight < b.totalWeight) {
      return 1;
    }
    return 0;
  });

  let maxTotalWeight: number = Object.values(activePrefs.prefs).reduce(
    (sum: number, userpref: UserPref): number => sum + userpref.priority,
    0,
  );

  if (sortedItems[0].totalWeight > maxTotalWeight) {
    maxTotalWeight = sortedItems[0].totalWeight;
  }

  const favs: FavRecord = getFav();

  return (
    <>
      <WhiteBlock className="w-full">
        <h1 className="text-1xl font-bold"> Получено {sortedItems.length} товаров</h1>
      </WhiteBlock>
      {sortedItems.map((item: ItemExtended, index: number): ReactElement => {
        const itemBaseData: Item = "item" in item ? item.item! : ({} as Item);
        const urls = [
          itemBaseData.url,
          ...item.similar.map((similarItem: Item): string => similarItem.url),
        ];
        const matchingUrl: string | undefined = urls.find((url: string): boolean =>
          Object.keys(favs).includes(url),
        );
        const itemData: ItemContextType = {
          item: itemBaseData,
          chars: item.chars,
          similar: item.similar,
          totalWeight: item.totalWeight,
          market: market,
          maxTotalWeight: maxTotalWeight,
          isFav: matchingUrl ? true : false,
          favUrl: matchingUrl ? matchingUrl : itemBaseData.url,
        };
        const filterData: FilterContextType = {
          categoryURL: categoryUrl,
          form: formPrefs ? formPrefs : ({} as PrefForm),
        };
        return (
          <ItemContextProvider
            data={itemData}
            key={index}
          >
            <FilterContextProvider data={filterData}>
              <ItemBlock />
            </FilterContextProvider>
          </ItemContextProvider>
        );
      })}
    </>
  );
}
