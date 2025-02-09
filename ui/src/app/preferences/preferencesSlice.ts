/* eslint-disable @typescript-eslint/no-use-before-define */
import { createSlice } from '@reduxjs/toolkit';
import { RootState } from '~/store';
import { Theme, Timezone } from '~/types/Preferences';

interface IPreferencesState {
  theme: Theme;
  timezone: Timezone;
}

const initialState: IPreferencesState = {
  theme: Theme.LIGHT,
  timezone: Timezone.LOCAL
};

export const preferencesSlice = createSlice({
  name: 'preferences',
  initialState,
  reducers: {
    themeChanged: (state, action) => {
      state.theme = action.payload;
    },
    timezoneChanged: (state, action) => {
      state.timezone = action.payload;
    }
  }
});

export const { themeChanged, timezoneChanged } = preferencesSlice.actions;

export const selectTheme = (state: RootState) => state.preferences.theme;
export const selectTimezone = (state: RootState) => state.preferences.timezone;

export default preferencesSlice.reducer;
