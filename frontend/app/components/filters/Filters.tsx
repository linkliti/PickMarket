import WhiteBlock from "@/components/base/WhiteBlock";
import FilterElem from "@/components/filters/FilterElem";

import { Filter } from "@/proto/app/protos/items";
import { LoadingSpinner } from "@/utilities/LoadingSpinner";
import { JsonValue } from "@protobuf-ts/runtime";
import axios, { AxiosResponse } from "axios";
import { ReactElement, useEffect, useState } from "react";
import terminal from "virtual:terminal";

export default function Filters({
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
    <WhiteBlock className="w-full">
      <>
        {isLoading ? (
          <div className="flex items-center gap-2">
            <LoadingSpinner /> <p>Загрузка фильтров</p>
          </div>
        ) : (
          <div>
            <h1 className="text-2xl font-bold"> Настройка фильтров:</h1>
            {filters.map((filter) => (
              <FilterElem
                filter={filter}
                key={filter.key}
              />
            ))}
          </div>
        )}
      </>
    </WhiteBlock>
  );
}
