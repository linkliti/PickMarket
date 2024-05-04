import FilterWrapper from "@/components/filters/filterTypes/FilterWrapper";
import MultipleSelector, { Option } from "@/components/ui/multiple-selector";
import { ReactElement } from "react";
import { Control, Controller } from "react-hook-form";

export default function SelectionPref({
  control,
  name,
  keyName,
  options,
}: {
  control: Control;
  name: string;
  keyName: string;
  options: Option[];
}): ReactElement {
  return (
    <FilterWrapper
      name={name}
      keyName={keyName}
    >
      <Controller
        name={keyName}
        control={control}
        defaultValue=""
        render={({ field }) => (
          <MultipleSelector
            className="mx-6 mt-2 bg-white"
            defaultOptions={options}
            placeholder="Выберите предпочтения..."
            emptyIndicator={
              <p className="text-center text-lg leading-10 text-gray-500 dark:text-gray-500">
                Не найдено
              </p>
            }
            {...field}
          />
        )}
      />
    </FilterWrapper>
  );
}
