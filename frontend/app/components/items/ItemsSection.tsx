import WhiteBlock from "@/components/base/WhiteBlock";
import ItemBlock from "@/components/items/item/ItemBlock";
import { ItemExtended, ItemsRequestWithPrefs, UserPref } from "@/proto/app/protos/reqHandlerTypes";
import { Markets } from "@/proto/app/protos/types";
import { FilterStore, useFilterStore } from "@/store/filterStore";
import { LoadingSpinner } from "@/utilities/LoadingSpinner";
import { JsonValue } from "@protobuf-ts/runtime";
import axios, { AxiosResponse } from "axios";
import { ReactElement, useEffect, useState } from "react";
import terminal from "virtual:terminal";

export default function ItemsSection({ market }: { market: number }): ReactElement {
  const [activePrefs] = useFilterStore((state: FilterStore) => [state.activePrefs]);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [items, setItems] = useState<ItemExtended[]>([]);
  const [error, setError] = useState<string>();

  // function dummyDataParser(): ItemExtended[] {
  //   const itemsDummy: ItemExtended[] = dummyData.map(
  //     (item: JsonValue): ItemExtended => ItemExtended.fromJson(item),
  //   );
  //   return itemsDummy;
  // }

  useEffect((): void => {
    async function getItems(): Promise<void> {
      try {
        if (activePrefs) {
          terminal.log("Posting items request");
          const res: AxiosResponse = await axios.post<JsonValue[]>(
            `/api/calc/${Markets[market].toLowerCase()}/list`,
            ItemsRequestWithPrefs.toJson(activePrefs),
            {
              headers: { "Content-Type": "application/json" },
              timeout: 120 * 1000,
            },
          );
          const itemsTemp: ItemExtended[] = res.data.map(
            (item: JsonValue): ItemExtended => ItemExtended.fromJson(item),
          );
          setItems(itemsTemp);
          setIsLoading(false);
        }
      } catch (error) {
        if (error instanceof Error) {
          terminal.error(error);
          setError(error.message);
        }
      }
    }

    // setItems(dummyDataParser());
    // setIsLoading(false);
    if (activePrefs) {
      setIsLoading(true);
      setItems([]);
      getItems();
    }
  }, [activePrefs, market]);

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
