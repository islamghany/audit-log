import { Text } from "../Typo";
interface Props {
  type: string;
  error?: string;
  name: string;
  label?: string;
  register: any;
}

export const Input = (props: Props) => {
  const { name, type, error, label, register } = props;
  return (
    <div className="form__unit">
      {label && (
        <label htmlFor={name}>
          <Text className="capitalize mb-2 block text-sm font-medium">
            {label}
          </Text>
        </label>
      )}
      <input
        type={type}
        name={name}
        {...register}
        className="block w-full appearance-none rounded-md border border-gray-200 bg-gray-50 px-3 py-2 text-gray-900 placeholder-gray-400 focus:border-primary focus:bg-white focus:outline-none focus:ring-primary sm:text-sm"
      />
      {error && <Text className="text-red-500">{error}</Text>}
    </div>
  );
};
