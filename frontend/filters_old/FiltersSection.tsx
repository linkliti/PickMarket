import WhiteBlock from "@/components/base/WhiteBlock";
import FilterSelectorMultiple from "@/components/filters/selectors/FilterSelectorMultiple";

import { Filter } from "@/proto/app/protos/items";
import { Markets } from "@/proto/app/protos/types";
import { blacklistKeys, useFiltersStore } from "@/store/filtersStore";
import { FiltersStore } from "@/types/filterTypes";
import { LoadingSpinner } from "@/utilities/LoadingSpinner";
import { JsonValue } from "@protobuf-ts/runtime";
import axios, { AxiosResponse } from "axios";
import { ReactElement, useEffect, useState } from "react";
import terminal from "virtual:terminal";

export default function FiltersSection({
  market,
  category,
}: {
  market: number;
  category: string;
}): ReactElement {
  const [filters, setFilters] = useState<Filter[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [resetStore, setPageData] = useFiltersStore((state: FiltersStore) => [
    state.resetStore,
    state.setPageData,
  ]);

  async function getFilters(market: number, category: string): Promise<void> {
    try {
      const res: AxiosResponse = await axios.get<JsonValue[]>(
        `/api/categories/${Markets[market].toLowerCase()}/filter?url=${category}`,
      );
      const filtersTemp: Filter[] = res.data
        .map((item: JsonValue): Filter => Filter.fromJson(item))
        .filter((item: Filter): boolean => !blacklistKeys.includes(item.key));
      setFilters(filtersTemp);
    } catch (error) {
      terminal.error(error);
    }
    return;
  }

  useEffect((): void => {
    // Store
    resetStore();
    setPageData(market, category);
    // Fetching
    getFilters(market, category);
    setIsLoading(false);
  }, [category, market, resetStore, setPageData]);

  return (
    <WhiteBlock className="w-full grow">
      <>
        {isLoading ? (
          <div className="flex items-center gap-2">
            <LoadingSpinner /> <p>Загрузка фильтров</p>
          </div>
        ) : (
          <>
            <h1 className="pb-4 text-2xl font-bold"> Настройка предпочтений:</h1>
            {filters.length === 0 ? (
              <p className="">Не удалось загрузить фильтры</p>
            ) : (
              <div className="grid grid-cols-1 gap-1 md:grid-cols-2 lg:grid-cols-3">
                <FilterSelectorMultiple filters={filters.slice(0, filters.length / 3)} />
                <FilterSelectorMultiple
                  filters={filters.slice(filters.length / 3, (2 * filters.length) / 3)}
                />
                <FilterSelectorMultiple filters={filters.slice((2 * filters.length) / 3)} />
              </div>
            )}
          </>
        )}
      </>
    </WhiteBlock>
  );
}
