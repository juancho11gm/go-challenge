import { shallowMount } from '@vue/test-utils'
import SearchDomain from '@/components/SearchDomain.vue'

describe('SearchDomain.vue', () => {
  it('renders props.msg when passed', () => {
    const msg = 'new message'
    const wrapper = shallowMount(SearchDomain, {
      propsData: { msg }
    })
    expect(wrapper.text()).toMatch(msg)
  })
})
