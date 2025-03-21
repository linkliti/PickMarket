import { useToast } from "@/components/ui/use-toast";
import { TriangleAlert } from "lucide-react";
import { ReactElement } from "react";
import { FieldErrors } from "react-hook-form";

export default function useErrorToast() {
  const { toast } = useToast();

  function errorToast(errors: FieldErrors): void {
    const toastContent = Object.values(errors)
      .map((error) => error?.message?.toString())
      .concat(Object.values(errors).map((error) => error?.root?.message?.toString()))
      .filter(Boolean) as string[];

    toast({
      className: "p-4 border border-border",
      action: (
        <>
          <TriangleAlert className="mr-2" />
          <div className="w-full items-center">
            {toastContent.length === 0 ? (
              <p className="first-letter:capitalize">Ошибка обработки формы</p>
            ) : (
              <>
                {toastContent.map(
                  (item: string, index: number): ReactElement => (
                    <p
                      className="first-letter:capitalize"
                      key={index}
                    >
                      {item}
                    </p>
                  ),
                )}
              </>
            )}
          </div>
        </>
      ),
      duration: 3000,
    });
  }
  return { errorToast };
}
