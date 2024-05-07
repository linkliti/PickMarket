import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
} from "@/components/ui/select";
import { PrefForm } from "@/types/filterTypes";
import { Star } from "lucide-react";
import { ReactElement } from "react";
import { Control, Controller } from "react-hook-form";

type PriorityItem = {
  value: string;
  label: string;
};

const prefItems: PriorityItem[] = [
  { value: "0", label: "Без разницы" },
  { value: "1", label: "Необязательно" },
  { value: "2", label: "Полезно" },
  { value: "3", label: "Важно" },
  { value: "4", label: "Очень важно" },
  { value: "5", label: "Обязательно" },
];
// Define the items as a constant

export default function PrioritySelector({
  control,
  keyName,
}: {
  control: Control<PrefForm, unknown>;
  keyName: string;
}): ReactElement {
  return (
    <Controller
      name={`priorities.${keyName}`}
      control={control}
      defaultValue={0}
      render={({ field: { onChange, value, disabled, name } }) => {
        function handleOnChange(selectedValue: string): void {
          const num: number = Number(selectedValue);
          onChange(num);
        }

        return (
          <Select
            onValueChange={handleOnChange}
            defaultValue={value.toString()}
            disabled={disabled}
            name={name}
          >
            <SelectTrigger className="h-8 w-fit px-2 py-0">
              {value}
              <Star className="ml-1 size-4" />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                {prefItems.map(
                  (item: PriorityItem): ReactElement => (
                    <SelectItem
                      key={item.value}
                      value={item.value.toString()}
                    >
                      {item.label}
                    </SelectItem>
                  ),
                )}
              </SelectGroup>
            </SelectContent>
          </Select>
        );
      }}
    />
  );
}
