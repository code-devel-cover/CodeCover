import Vue from 'vue';
import { createLocalVue, shallowMount, Wrapper } from '@vue/test-utils';
import Vuex, { Store } from 'vuex';
import { cloneDeep } from 'lodash';
import Vuetify from 'vuetify';
import flushPromises from 'flush-promises';
import ReportOverview from '@/components/ReportOverview.vue';
import { AxiosPlugin } from '@/plugins/http';
import { Mutations, storeConfig, RootState } from '@/store';

jest.mock('axios');
Vue.use(Vuetify);
Vue.use(AxiosPlugin);

describe('ReportOverview.vue', () => {
  console.warn = jest.fn();
  const localVue = createLocalVue();
  let vuetify: typeof Vuetify;
  let store: Store<RootState>;
  beforeEach(() => {
    vuetify = new Vuetify();
    store = new Vuex.Store(cloneDeep(storeConfig));
  });

  localVue.use(Vuetify);
  it('update file count when current report is set', async () => {
    const wrapper = shallowMount(ReportOverview, {
      localVue,
      vuetify,
      store
    }) as Wrapper<ReportOverview & { filesCount: number }>;
    expect(wrapper.vm.$store.state.report.current).toBeUndefined();
    expect(wrapper.vm.filesCount).toEqual(0);
    wrapper.vm.$store.commit(Mutations.SET_REPORT_CURRENT, {
      files: ['a', 'b', 'c']
    } as Report);
    await flushPromises();
    expect(wrapper.vm.filesCount).toEqual(3);
  });

  it('round coverage rate', async () => {
    const wrapper = shallowMount(ReportOverview, {
      localVue,
      vuetify,
      store
    }) as Wrapper<ReportOverview & { $coverage: number }>;
    expect(wrapper.vm.$store.state.report.current).toBeUndefined();
    expect(wrapper.vm.$coverage).toEqual(0);
    wrapper.vm.$store.commit(Mutations.SET_REPORT_CURRENT, {
      coverages: [
        {
          statementCoverage: 0.8995
        }
      ]
    } as Report);
    await flushPromises();
    expect(wrapper.vm.$coverage).toEqual(89.95);
  });
});
