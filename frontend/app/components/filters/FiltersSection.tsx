import WhiteBlock from "@/components/base/WhiteBlock";
import FilterForm from "@/components/filters/FilterForm";

import { Button } from "@/components/ui/button";
import { Option } from "@/components/ui/multiple-selector";
import { Filter } from "@/proto/app/protos/items";
import { ItemsRequestWithPrefs } from "@/proto/app/protos/reqHandlerTypes";
import { Markets } from "@/proto/app/protos/types";

import { blacklistKeys } from "@/store/filtersStore";
import { LoadingSpinner } from "@/utilities/LoadingSpinner";
import { JsonValue } from "@protobuf-ts/runtime";
import axios, { AxiosResponse } from "axios";
import { ReactElement, useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import terminal from "virtual:terminal";

type PrefForm = {
  [key in string]: number | boolean | Option[];
};

export default function FiltersSection({
  market,
  category,
}: {
  market: number;
  category: string;
}): ReactElement {
  const [isLoading, setIsLoading] = useState(true);
  const [filters, setFilters] = useState<Filter[]>([]);
  const { handleSubmit, control: formControl } = useForm<PrefForm>();

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
  }

  function onSubmit(data: PrefForm): void {
    const req: ItemsRequestWithPrefs = {
      request: {
        market: market,
        pageUrl: category,
        numOfPages: 1,
        params: "",
        userQuery: "",
      },
      prefs: {},
    };

    Object.entries(data).forEach(([key, value]: [string, number | boolean | Option[]]): void => {
      terminal.log(value);
      switch (typeof value) {
        case "number": {
          if (!value) break;
          req.prefs[key] = { priority: 0, value: { oneofKind: "numVal", numVal: value } };
          break;
        }
        case "boolean": {
          req.prefs[key] = {
            priority: 0,
            value: {
              oneofKind: "listVal",
              listVal: {
                values: value ? ["Да"] : ["Нет"],
              },
            },
          };
          break;
        }
        case "object": {
          if (!value.length) break; // value is an array with length of at least 1
          req.prefs[key] = {
            priority: 0,
            value: {
              oneofKind: "listVal",
              listVal: {
                values: value.map((item: Option): string => item.value),
              },
            },
          };
        }
      }
    });

    console.log(req);
  }

  useEffect((): void => {
    // Fetching
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
          <form onSubmit={handleSubmit(onSubmit)}>
            <h1 className="pb-4 text-2xl font-bold"> Настройка предпочтений:</h1>
            <div className="grid grid-cols-1 gap-1 md:grid-cols-2 lg:grid-cols-3">
              {[
                filters.slice(0, filters.length / 3),
                filters.slice(filters.length / 3, (2 * filters.length) / 3),
                filters.slice((2 * filters.length) / 3),
              ].map(
                (group: Filter[], index: number): ReactElement => (
                  <div key={index}>
                    {group.map(
                      (filter: Filter): ReactElement => (
                        <FilterForm
                          control={formControl}
                          filter={filter}
                        />
                      ),
                    )}
                  </div>
                ),
              )}
            </div>
            <Button type="submit">Применить</Button>
          </form>
        )}
      </>
    </WhiteBlock>
  );
}
