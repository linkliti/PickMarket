import WhiteBlock from "@/components/base/WhiteBlock";
import FilterSelector from "@/components/filters/selectors/FilterSelector";

import { Filter } from "@/proto/app/protos/items";
import { LoadingSpinner } from "@/utilities/LoadingSpinner";
import { JsonValue } from "@protobuf-ts/runtime";
import axios, { AxiosResponse } from "axios";
import { ReactElement, useEffect, useState } from "react";
import terminal from "virtual:terminal";

// const blacklistKeys = ["pm_isAdult", "trucode"];

export default function FiltersSection({
  market,
  category,
}: {
  market: string;
  category: string;
}): ReactElement {
  const [filters, setFilters] = useState<Filter[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  async function getFilters(market: string, category: string): Promise<void> {
    try {
      const res: AxiosResponse = await axios.get<JsonValue[]>(
        `/api/categories/${market}/filter?url=${category}`,
      );
      const data = res.data;
      const filtersTemp: Filter[] = [];
      data.map((item: JsonValue): void => {
        filtersTemp.push(Filter.fromJson(item));
      });
      setFilters(filtersTemp);
    } catch (error) {
      terminal.error(error);
    }
    return;
  }

  useEffect((): void => {
    getFilters(market, category);
    setIsLoading(false);
  }, [category, market]);

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
            <div className="grid grid-cols-1 gap-1 md:grid-cols-2 lg:grid-cols-3">
              <FilterSelectorMultiple filters={filters.slice(0, filters.length / 3)} />
              <FilterSelectorMultiple
                filters={filters.slice(filters.length / 3, (2 * filters.length) / 3)}
              />
              <FilterSelectorMultiple filters={filters.slice((2 * filters.length) / 3)} />
            </div>
          </>
        )}
      </>
    </WhiteBlock>
  );
}

function FilterSelectorMultiple({ filters }: { filters: Filter[] }): ReactElement {
  return (
    <div className="col-span-1">
      {filters.map((filter) => (
        <FilterSelector
          filter={filter}
          key={filter.key}
          className="rounded border-b border-b-gray-200"
        />
      ))}
    </div>
  );
}
