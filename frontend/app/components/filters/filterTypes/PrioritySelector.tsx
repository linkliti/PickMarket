import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { PrefForm } from "@/types/filterTypes";
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
        const handleOnChange = (selectedValue: string) => {
          const num = Number(selectedValue);
          onChange(num);
        };

        return (
          <Select
            onValueChange={handleOnChange}
            defaultValue={value.toString()}
            disabled={disabled}
            name={name}
          >
            <SelectTrigger className="h-8 w-fit px-2 py-0">
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              <SelectGroup>
                {prefItems.map((item) => (
                  <SelectItem
                    key={item.value}
                    value={item.value.toString()}
                  >
                    {item.label}
                  </SelectItem>
                ))}
              </SelectGroup>
            </SelectContent>
          </Select>
        );
      }}
    />
  );
}

// function PriorityCombobox({
//   className = "",
//   listClassNames = "",
// }: {
//   className?: string;
//   listClassNames?: string;
// }): ReactElement {
//   const [isOpen, setOpen] = useState<boolean>(false);
//   const [selected, setSelected] = useState<string>("0");
//
//   return (
//     <div className={cn(className)}>
//       <Popover
//         open={isOpen}
//         onOpenChange={setOpen}
//       >
//         <PopoverTrigger asChild>
//           <Button
//             variant="outline"
//             role="combobox"
//             aria-expanded={isOpen}
//             className={cn("justify-between", listClassNames)}
//           >
//             {selected ? selected : "Выбрать..."}
//             <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
//           </Button>
//         </PopoverTrigger>
//         <PopoverContent className={cn("p-0", listClassNames)}>
//           <Command>
//             {/* <CommandInput placeholder="Выбрать маркетплейс..." /> */}
//             <CommandList>
//               <CommandEmpty>Маркетплейс не найден</CommandEmpty>
//               <CommandGroup>
//                 {prefItems.map(
//                   (item: PriorityItem): ReactElement => (
//                     <CommandItem
//                       key={item.value}
//                       value={item.value}
//                       onSelect={(currentValue: string): void => {
//                         setSelected(currentValue);
//                         setOpen(false);
//                       }}
//                     >
//                       <Check
//                         className={cn(
//                           "mr-2 h-4 w-4",
//                           item.value === selected ? "opacity-100" : "opacity-0",
//                         )}
//                       />
//                       {item.label}
//                     </CommandItem>
//                   ),
//                 )}
//               </CommandGroup>
//             </CommandList>
//           </Command>
//         </PopoverContent>
//       </Popover>
//     </div>
//   );
// }
