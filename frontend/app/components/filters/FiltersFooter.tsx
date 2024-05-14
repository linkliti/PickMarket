import { Button } from "@/components/ui/button";
import { CollapsibleTrigger } from "@/components/ui/collapsible";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { FilterStore, useFilterStore } from "@/store/filterStore";
import { PrefForm } from "@/types/filterTypes";
import { ChevronsUpDown } from "lucide-react";
import { ReactElement } from "react";
import { Control, Controller, UseFormReset } from "react-hook-form";

export default function FiltersFooter({
  isOpen,
  control,
  reset,
  isSubmitting,
}: {
  isOpen: boolean;
  control: Control<PrefForm, unknown>;
  reset: UseFormReset<PrefForm>;
  isSubmitting: boolean;
}) {
  const [setFormPrefs] = useFilterStore((state: FilterStore) => [state.setFormPrefs]);

  return (
    <div className="inline-flex w-full items-end gap-2 max-sm:flex-wrap">
      <div className="grow">
        <Label>Поисковой запрос</Label>
        <Controller
          control={control}
          name="userQuery"
          defaultValue=""
          render={({ field }): ReactElement => {
            return (
              <Input
                className="bg-white"
                {...field}
              />
            );
          }}
        />
      </div>
      <div className="w-[100px] grow">
        <Label>Страниц</Label>
        <Controller
          control={control}
          name="numOfPages"
          defaultValue={1}
          render={({ field: { onChange, onBlur, value, disabled, name, ref } }) => {
            return (
              <Input
                type="number"
                min={1}
                max={5}
                step={1}
                className=" bg-white"
                onChange={(event: React.ChangeEvent<HTMLInputElement>): void => {
                  const num: number = parseFloat(event.target.value);
                  if (isNaN(num) || num === 0) {
                    onChange(0);
                    return;
                  }
                  if (num > 5) {
                    onChange(5);
                    return;
                  }
                  onChange(num);
                }}
                onBlur={onBlur}
                value={Number(value).toString()}
                disabled={disabled}
                name={name}
                ref={ref}
              />
            );
          }}
        />
      </div>
      <div className="inline-flex gap-2 max-sm:flex-wrap">
        <Button
          type="submit"
          disabled={isSubmitting}
        >
          Применить
        </Button>
        <Button
          onClick={(event): void => {
            event.preventDefault();
            setFormPrefs({} as PrefForm);
            reset({});
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
  );
}
