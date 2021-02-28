import {useCallback} from 'react'

export const useMessage = () => {
  return useCallback(text => {
    if (window.M && text) { //function M is from materialized view (checks if M exists and text isDefined)
      window.M.toast({ html: text })
    }
  }, [])
}