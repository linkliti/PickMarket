import WhiteBlock from "@/components/base/WhiteBlock";
import ItemBlock from "@/components/items/item/ItemBlock";
import useFetchData from "@/components/items/useFetchItems";
import { ItemExtended, UserPref } from "@/proto/app/protos/reqHandlerTypes";
import { FilterStore, useFilterStore } from "@/store/filterStore";
import { LoadingSpinner } from "@/utilities/LoadingSpinner";

import { ReactElement } from "react";

export default function ItemsSection(): ReactElement {
  const [activePrefs, market] = useFilterStore((state: FilterStore) => [
    state.activePrefs,
    state.market,
  ]);
  const { items, isLoading, error } = useFetchData();

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

  return (
    <>
      <WhiteBlock className="w-full">
        <h1 className="text-1xl font-bold"> Получено {sortedItems.length} товаров</h1>
      </WhiteBlock>
      {sortedItems.map(
        (item: ItemExtended, index: number): ReactElement => (
          <ItemBlock
            market={market}
            key={index}
            item={item}
            maxTotalWeight={maxTotalWeight}
          />
        ),
      )}
    </>
  );
}
