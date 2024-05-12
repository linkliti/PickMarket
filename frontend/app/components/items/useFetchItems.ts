import { ItemExtended, ItemsRequestWithPrefs } from "@/proto/app/protos/reqHandlerTypes";
import { Markets } from "@/proto/app/protos/types";
import { FilterStore, useFilterStore } from "@/store/filterStore";
import { JsonValue } from "@protobuf-ts/runtime";
import axios, { AxiosResponse } from "axios";
import { useEffect, useState } from "react";
import terminal from "virtual:terminal";

export default function useFetchItems() {
  const [activePrefs, market] = useFilterStore((state: FilterStore) => [state.activePrefs, state.market]);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [items, setItems] = useState<ItemExtended[]>([]);
  const [error, setError] = useState<string>();

  useEffect((): void => {
    async function getItems(): Promise<void> {
      try {
        setError(undefined);
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

  return { items, isLoading, error, setItems };
}