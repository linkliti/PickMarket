import WhiteBlock from "@/components/base/WhiteBlock";
import FilterForm from "@/components/filters/FilterForm";

import { Button } from "@/components/ui/button";
import { Filter } from "@/proto/app/protos/items";
import { ItemsRequestWithPrefs } from "@/proto/app/protos/reqHandlerTypes";
import { Markets } from "@/proto/app/protos/types";
import { PrefForm } from "@/types/filterTypes";
import { JsonValue } from "@protobuf-ts/runtime";

import { Collapsible, CollapsibleContent, CollapsibleTrigger } from "@/components/ui/collapsible";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { cn } from "@/lib/utils";
import { FilterStore, blacklistKeys, useFilterStore } from "@/store/filterStore";
import { LoadingSpinner } from "@/utilities/LoadingSpinner";
import { useQuery } from "@tanstack/react-query";
import axios, { AxiosResponse } from "axios";
import { ChevronsUpDown } from "lucide-react";
import { ReactElement, useState } from "react";
import { Controller, useForm } from "react-hook-form";
import terminal from "virtual:terminal";

export default function FiltersSection({
  market,
  category,
}: {
  market: number;
  category: string;
}): ReactElement {
  const [isOpen, setIsOpen] = useState<boolean>(true);
  const [formPrefs, setActivePrefs, setFormPrefs] = useFilterStore((state: FilterStore) => [
    state.formPrefs,
    state.setActivePrefs,
    state.setFormPrefs,
  ]);
  const { handleSubmit, control, reset } = useForm<PrefForm>({
    defaultValues: formPrefs ? formPrefs : {},
  });

  const {
    isPending: isLoading,
    error,
    data: filters,
  } = useQuery({
    queryKey: ["filters", market, category],
    queryFn: getFilters,
    staleTime: Infinity,
  });

  function onSubmit(data: PrefForm): void {
    const req: ItemsRequestWithPrefs = {
      request: {
        market: market,
        pageUrl: category,
        numOfPages: data.numOfPages,
        params: "",
        userQuery: data.userQuery || "",
      },
      prefs: {},
    };

    for (const [key, priority] of Object.entries(data.priorities)) {
      if (!priority) continue;
      const charData: number | boolean | string[] = data.prefs[key];
      switch (typeof charData) {
        case "number": {
          req.prefs[key] = { priority: priority, value: { oneofKind: "numVal", numVal: charData } };
          break;
        }
        case "boolean": {
          req.prefs[key] = {
            priority: priority,
            value: {
              oneofKind: "listVal",
              listVal: {
                values: charData ? ["Да"] : ["Нет"],
              },
            },
          };
          break;
        }
        case "object": {
          if (!charData.length) break;
          req.prefs[key] = {
            priority: priority,
            value: {
              oneofKind: "listVal",
              listVal: {
                values: charData,
              },
            },
          };
        }
      }
    }
    setActivePrefs(req);
    setIsOpen(false);
  }

  async function getFilters(): Promise<Filter[] | undefined> {
    try {
      terminal.log("Fetching", category);
      const res: AxiosResponse = await axios.get<JsonValue[]>(
        `/api/categories/${Markets[market].toLowerCase()}/filter?url=${category}`,
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

  return (
    <WhiteBlock className={cn("w-full")}>
      <Collapsible
        open={isOpen}
        onOpenChange={setIsOpen}
      >
        {isLoading ? (
          <div className="flex items-center gap-2">
            <LoadingSpinner /> <p>Загрузка фильтров</p>
          </div>
        ) : !filters || error ? (
          <p>Ошибка при загрузке фильтров: {error?.message}</p>
        ) : (
          <>
            <form onSubmit={handleSubmit(onSubmit)}>
              <h1 className="pb-4 text-2xl font-bold"> Настройка предпочтений:</h1>
              <CollapsibleContent>
                <div className="grid grid-cols-1 gap-1 pb-4 md:grid-cols-2 lg:grid-cols-3">
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
                              control={control}
                              filter={filter}
                              key={filter.key}
                            />
                          ),
                        )}
                      </div>
                    ),
                  )}
                </div>
              </CollapsibleContent>
              <div className="inline-flex w-full flex-wrap items-end gap-2">
                <div className="grow">
                  <Label>Поисковой запрос</Label>
                  <Controller
                    control={control}
                    name="userQuery"
                    defaultValue=""
                    render={({ field }): ReactElement => {
                      return (
                        <Input
                          className=" max-w-[300px] bg-white"
                          {...field}
                        />
                      );
                    }}
                  ></Controller>
                </div>
                <div className="">
                  <Label>Страниц</Label>
                  <Controller
                    control={control}
                    name="numOfPages"
                    defaultValue={1}
                    render={({ field: { onChange, onBlur, value, disabled, name, ref } }) => {
                      return (
                        <Input
                          className=" bg-white"
                          onChange={(event: React.ChangeEvent<HTMLInputElement>): void => {
                            const num: number = parseInt(event.target.value, 10);
                            if (isNaN(num) || num === 0) {
                              onChange(0);
                            } else {
                              onChange(num);
                            }
                          }}
                          onBlur={onBlur}
                          value={value}
                          disabled={disabled}
                          name={name}
                          ref={ref}
                        />
                      );
                    }}
                  ></Controller>
                </div>
                <div className="inline-flex flex-wrap gap-2">
                  <Button type="submit">Применить</Button>
                  <Button
                    onClick={(event): void => {
                      event.preventDefault();
                      setFormPrefs(null);
                      reset();
                    }}
                  >
                    Сбросить
                  </Button>
                  <CollapsibleTrigger asChild>
                    <Button>
                      {isOpen ? "Скрыть" : "Показать"}
                      <>
                        <ChevronsUpDown className="ml-1 h-4 w-4" />
                      </>
                    </Button>
                  </CollapsibleTrigger>
                </div>
              </div>
            </form>
          </>
        )}
      </Collapsible>
    </WhiteBlock>
  );
}
