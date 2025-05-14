import { useForm, FormProvider } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import { RegisterSchema, type TRegisterForm } from '@app/validation/auth';
import { Button } from '@app/components/ui/button';
import { Card, CardContent, CardTitle } from '@app/components/ui/card';
import RHFInput from '@app/components/hook-form/rhf-input';
import RHFSelect from '@app/components/hook-form/rhf-select';
import store, { observer } from '@app/stores';
import { EyeIcon, EyeOff, TriangleAlert } from 'lucide-react';

import { toast } from 'sonner';
import { postRegister } from '@app/api/auth';
import { useNavigate } from 'react-router';
import { errorHandler } from '@app/lib/axios';
import { Link } from 'react-router-dom';
import { useState } from 'react';

const countryOptions = [
  { label: '🇬🇧 +44', value: '+44' },
  { label: '🇺🇸 +1', value: '+1' },
  { label: '🇫🇷 +33', value: '+33' },
  { label: '🇩🇪 +49', value: '+49' },
  { label: '🇮🇳 +91', value: '+91' }
];

const RegisterForm = observer(() => {
  const [showPassword, setShowPassword] = useState(false);

  const methods = useForm<TRegisterForm>({
    resolver: yupResolver(RegisterSchema),
    defaultValues: {
      first_name: '',
      last_name: '',
      email: '',
      mobile_country: countryOptions.find((opt) => opt.value === '+44')?.value,
      mobile_number: '',
      password: '',
      confirm_password: ''
    }
  });

  const nav = useNavigate();

  const onSubmit = async (data: TRegisterForm) => {
    try {
      store.ui.setIsLoading(true);
      const res = await postRegister(data);
      store.auth.authenticate(res.data);
      nav('/');
    } catch (error) {
      errorHandler(error, (e) =>
        toast(e.message, {
          position: 'bottom-left',
          icon: <TriangleAlert className="text-destructive mr-10" />,
          description:
            'Please try again, or contact support if further help is needed.'
        })
      );
      store.auth.unauthenticate();
    } finally {
      store.ui.setIsLoading(false);
    }
  };

  return (
    <FormProvider {...methods}>
      <Card>
        <CardContent>
          <CardTitle className="mb-2">Create an Account</CardTitle>
          <div className="text-muted-foreground *:[a]:hover:text-primary text-left text-xs text-balance *:[a]:underline *:[a]:underline-offset-4">
            Already have an account? <Link to="/sign-in">Sign in</Link>
          </div>
          <form
            onSubmit={methods.handleSubmit(onSubmit)}
            className="mt-6 flex flex-col gap-5"
          >
            <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
              <RHFInput
                placeholder="John"
                name="first_name"
                label="First Name"
              />
              <RHFInput placeholder="Doe" name="last_name" label="Last Name" />
            </div>
            <RHFInput
              placeholder="your@email.com"
              name="email"
              label="Email"
              type="email"
            />
            <div className="grid grid-cols-[auto_1fr] gap-4">
              <RHFSelect
                name="mobile_country"
                label="Country Code"
                options={countryOptions}
                placeholder="Code"
              />
              <RHFInput
                name="mobile_number"
                label="Mobile Number"
                placeholder="7123456789"
              />
            </div>
            <RHFInput
              name="password"
              label="Password"
              type={showPassword ? 'text' : 'password'}
              endIcon={
                showPassword ? (
                  <EyeIcon
                    className="size-4"
                    onClick={() => setShowPassword(!showPassword)}
                  />
                ) : (
                  <EyeOff
                    className="size-4"
                    onClick={() => setShowPassword(!showPassword)}
                  />
                )
              }
            />
            <RHFInput
              name="confirm_password"
              label="Confirm Password"
              type={showPassword ? 'text' : 'password'}
            />
            <Button type="submit" className="w-full">
              Register
            </Button>
            <div className="text-muted-foreground *:[a]:hover:text-primary text-center text-xs text-balance *:[a]:underline *:[a]:underline-offset-4">
              By signing up, you agree to our <a href="#">Terms of Service</a>{' '}
              and <a href="#">Privacy Policy</a>.
            </div>
          </form>
        </CardContent>
      </Card>
    </FormProvider>
  );
});

export default RegisterForm;
