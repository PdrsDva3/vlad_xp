import './styles/App.scss';
import { useState, useEffect, useRef } from 'react';
import { Input, Info } from './components';
import { checkHasError, checkDataMakeError, checkData, sendFormData } from './utils';
import { EMPTY_FORM_DATA, EMPTY_FORM_ERROR, REGULAR_EMAIL } from './constants';
import axios from 'axios';

function App() {
	const [formData, setFormData] = useState(EMPTY_FORM_DATA);
	const [formError, setFormError] = useState(EMPTY_FORM_ERROR);
	const [hasError, setHasError] = useState(false);

	const submitButtonRef = useRef(null);

	/// ↓ API ↓
	const USER_URL = 'http://localhost:63342/user';

	const getData = (id) => {
		axios.get(`${USER_URL}/get/${id}`).then((rawResponse) => {
			const response = rawResponse.data;
			console.log('response', response);
		});
	};

	// const postData = (formData) => {
	// 	fetch(`${DATA_URL}/create`, {
	// 		method: "POST",
	// 		headers: { "Content-Type": "application/json;charset=utf-8" },
	// 		body: JSON.stringify({
	// 			// title: title,
	// 		}),
	// 	})
	// 	.then((res) => {
	// 		if (res.ok) {
	// 			return res.data;
	// 		}
	// 	})
	// 	.catch((e) => {
	// 		throw e;
	// 	});
	// }

	// useEffect(() => {
	// 	getData();
	// }, []);

	/// ↑ API ↑

	useEffect(() => {
		setHasError(checkHasError(formError));
	}, [formError]);

	// EMAIL CHANGE
	const onEmailChange = ({ target }) => {
		setFormData({ ...formData, email: target.value });
		setFormError({ ...formError, emailError: null });
	};

	// EMAIL BLUR
	const onEmailBlur = ({ target }) => {
		let newEmailError = null;

		if (!REGULAR_EMAIL.test(target.value)) {
			newEmailError = 'Invalid email.';
		}
		if (target.value.length === 0) {
			newEmailError = "Field shouldn't be empty";
		}

		setFormError({ ...formError, emailError: newEmailError });
	};

	// PASSWORD CHANGE
	const onPasswordChange = ({ target }) => {
		setFormData({ ...formData, password: target.value });

		let newPasswordError = null;

		if (target.value.length > 20) {
			newPasswordError = 'No more than 20 chars';
		}

		setFormError({
			...formError,
			passwordError: newPasswordError,
			rePasswordError: null,
		});
	};

	// PASSWORD BLUR
	const onPasswordBlur = ({ target }) => {
		let newPasswordError = formError.passwordError;

		if (target.value.length === 0) {
			newPasswordError = "Field shouldn't be empty";
		} else if (formData.rePassword && target.value !== formData.rePassword) {
			newPasswordError = "Passwords don't match.";
		}

		setFormError({
			...formError,
			passwordError: newPasswordError,
		});
	};

	// SECOND PASSWORD CHANGE
	const onRePasswordChange = ({ target }) => {
		setFormData({ ...formData, rePassword: target.value });

		let newRePasswordError = null;

		if (
			target.value === formData.password &&
			!checkHasError(formError) &&
			formData.email
		) {
			submitButtonRef.current.focus();
		}

		setFormError({
			...formError,
			passwordError: null,
			rePasswordError: newRePasswordError,
		});
	};

	// SECOND PASSWORD BLUR
	const onRePasswordBlur = ({ target }) => {
		let newRePasswordError = formError.rePasswordError;

		if (target.value.length === 0) {
			newRePasswordError = "Field shouldn't be empty";
		} else if (target.value !== formData.password) {
			newRePasswordError = "Passwords don't match.";
		}
		setFormError({ ...formError, rePasswordError: newRePasswordError });
	};

	const postData = (formData) => {
		fetch(`${USER_URL}/create`, {
			method: "POST",
			headers: { "Content-Type": "application/json;charset=utf-8" },
			body: JSON.stringify({
				email: formData.email,
				password: formData.password,
			}),
		})
			.then((res) => {
				if (res.ok) {
					return res.data;
				}
			})
			.catch((e) => {
				throw e;
			});
	}

	// SUBMIT
	const onSubmit = (event) => {
		event.preventDefault();

		const newFormError = { ...formError };

		if (checkData(formData)) {
			sendFormData(formData);
			postData(formData)
			setFormData(EMPTY_FORM_DATA);
		} else {
			checkDataMakeError(formData, newFormError);
			setFormError({ ...newFormError });
		}
	};


	return (
		<section className="login">
			<div className="login__wrapper">
				<Info />
				<form className="form" onSubmit={onSubmit}>
					<h2 className="form__title">Sign up</h2>
					<Input
						type="email"
						name="email"
						onChange={onEmailChange}
						onBlur={onEmailBlur}
						value={formData.email}
						errorMessage={formError.emailError}
					/>
					<Input
						type="password"
						name="password"
						onChange={onPasswordChange}
						onBlur={onPasswordBlur}
						value={formData.password}
						errorMessage={formError.passwordError}
					/>
					<Input
						type="password"
						name="rePassword"
						label="Re-enter password"
						onChange={onRePasswordChange}
						onBlur={onRePasswordBlur}
						value={formData.rePassword}
						errorMessage={formError.rePasswordError}
					/>
					<span>
						<button
							className="form__submit-button"
							type="submit"
							ref={submitButtonRef}
							disabled={hasError}
						>
							Create account
						</button>
					</span>
				</form>
			</div>
		</section>
	);
}

export default App;
