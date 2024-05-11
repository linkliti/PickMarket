import { Filter } from "@/proto/app/protos/items";
import { Markets } from "@/proto/app/protos/types";
import { FilterStore, blacklistKeys, useFilterStore } from "@/store/filterStore";
import { JsonValue } from "@protobuf-ts/runtime";
import { useQuery } from "@tanstack/react-query";
import axios, { AxiosResponse } from "axios";

import terminal from "virtual:terminal";

export default function useFetchFilters() {

  const [market, categoryUrl] = useFilterStore(
    (state: FilterStore) => [

      state.market,
      state.categoryUrl,
    ],
  );

  const {
    isPending: isLoading,
    error,
    data: filters,
  } = useQuery({
    queryKey: ["filters", market, categoryUrl],
    queryFn: getFilters,
    staleTime: Infinity,
  });

  async function getFilters(): Promise<Filter[] | undefined> {
    try {
      terminal.log("Fetching", categoryUrl);
      const res: AxiosResponse = await axios.get<JsonValue[]>(
        `/api/categories/${Markets[market].toLowerCase()}/filter?url=${categoryUrl}`,
      );
      const filtersTemp: Filter[] = res.data
        .map((item: JsonValue): Filter => Filter.fromJson(item))
        .filter((item: Filter): boolean => !blacklistKeys.includes(item.key));
      return filtersTemp;
    } catch (error) {
      if (error instanceof Error) {
        terminal.error(error.message);
        throw new Error(error.message);
      }
    }
  }

  return { isLoading, error, filters };
}
