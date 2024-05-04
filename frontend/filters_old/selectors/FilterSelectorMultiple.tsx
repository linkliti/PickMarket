import FilterSelector from "@/components/filters/selectors/FilterSelector";
import { Filter } from "@/proto/app/protos/items";

import { ReactElement } from "react";

export default function FilterSelectorMultiple({ filters }: { filters: Filter[] }): ReactElement {
  return (
    <div className="col-span-1">
      {filters.map(
        (filter: Filter): ReactElement => (
          <FilterSelector
            filter={filter}
            key={filter.key}
            className="rounded border-b border-b-gray-200"
          />
        ),
      )}
    </div>
  );
}
