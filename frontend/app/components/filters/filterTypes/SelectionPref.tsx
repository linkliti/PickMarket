import FilterWrapper from "@/components/filters/filterTypes/FilterWrapper";
import MultipleSelector, { Option } from "@/components/ui/multiple-selector";
import { PrefForm } from "@/types/filterTypes";
import { ReactElement } from "react";
import { Control, Controller } from "react-hook-form";

export default function SelectionPref({
  control,
  filterTitle,
  keyName,
  options,
}: {
  control: Control<PrefForm, unknown>;
  filterTitle: string;
  keyName: string;
  options: Option[];
}): ReactElement {
  return (
    <Controller
      name={`prefs.${keyName}`}
      control={control}
      defaultValue={[]}
      render={({ field: { onChange, value, disabled, ref } }) => {
        const transformedValue: Option[] | undefined = Array.isArray(value)
          ? value.map((str): Option => ({ value: str, label: str }))
          : undefined;

        function handleOnChange(selectedOptions: Option[]): void {
          const stringArray: string[] = selectedOptions.map(
            (option: Option): string => option.value,
          );
          onChange(stringArray);
        }

        return (
          <FilterWrapper
            name={filterTitle}
            keyName={keyName}
            control={control}
          >
            <MultipleSelector
              className="mx-6 mt-2 bg-white"
              defaultOptions={options}
              placeholder="Выберите предпочтения..."
              hidePlaceholderWhenSelected={true}
              emptyIndicator={
                <p className="text-center text-lg leading-10 text-gray-500 dark:text-gray-500">
                  Не найдено
                </p>
              }
              onChange={handleOnChange}
              // onBlur={onBlur}
              value={transformedValue}
              disabled={disabled}
              // name={name}
              ref={ref}
            />
          </FilterWrapper>
        );
      }}
    />
  );
}
