import WhiteBlock from "@/components/base/WhiteBlock";
import FilterForm from "@/components/filters/FilterForm";

import { Filter } from "@/proto/app/protos/items";

import Loading from "@/components/base/Loading";
import FiltersFooter from "@/components/filters/FiltersFooter";
import useErrorToast from "@/components/filters/useErrorToast";
import useFetchFilters from "@/components/filters/useFetchFilters";
import useFilterForm from "@/components/filters/useFilterForm";
import { Collapsible, CollapsibleContent } from "@/components/ui/collapsible";
import { cn } from "@/lib/utils";
import { ReactElement } from "react";

export default function FiltersSection(): ReactElement {
  const { errorToast } = useErrorToast();
  const { isLoading, error, filters } = useFetchFilters();

  const { isOpen, handleSubmit, control, reset, onSubmit, isSubmitting, setIsOpen } =
    useFilterForm();

  return (
    <WhiteBlock className={cn("w-full flex-col")}>
      <Collapsible
        open={isOpen}
        onOpenChange={setIsOpen}
      >
        {isLoading ? (
          <Loading message="Загрузка предпочтений" />
        ) : !filters || error ? (
          <p>Ошибка при загрузке предпочтений: {error?.message}</p>
        ) : (
          <>
            <form onSubmit={handleSubmit(onSubmit, errorToast)}>
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
              <FiltersFooter
                isOpen={isOpen}
                control={control}
                reset={reset}
                isSubmitting={isSubmitting}
              />
            </form>
          </>
        )}
      </Collapsible>
    </WhiteBlock>
  );
}
