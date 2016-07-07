# coding: utf-8

from flask import url_for
from flask_admin.contrib.sqla import ModelView
from flask_admin.base import expose
from jinja2 import contextfunction


class PuzeApplyRecordAdmin(ModelView):

    # Views
    @expose('/')
    def index_view(self):
        """
            List view
        """

        # Grab parameters from URL
        view_args = self._get_list_extra_args()

        # 找出过滤信息, 自己定制的filter信息
        filters = view_args.filters
        args = {}
        """
        filters:
            [(0, <filter model>, u'filter_msg'), (), ...]
        """
        for item in filters:
            args[item[1].column.name] = item[2]

        # 自己实现
        data, num_pages = ([], 1)
        count = len(data)

        num_pages = num_pages if num_pages > 1 else None

        # Actions
        actions, actions_confirmation = self.get_actions_list()

        def pager_url(p):
            # Do not add page number if it is first page
            if p == 0:
                p = None

            return self._get_list_url(view_args.clone(page=p))

        this_url = url_for('.index_view')

        return self.render(
            self.list_template,
            data=data,
            form=None,
            delete_form=None,

            # List
            list_columns=self._list_columns,
            sortable_columns={},
            editable_columns=None,

            # Pagination
            count=count,
            pager_url=pager_url,
            num_pages=num_pages,
            page=view_args.page,
            page_size=self.page_size,

            # Sorting
            sort_column=None,
            sort_desc=view_args.sort_desc,
            sort_url=None,

            # Search
            search_supported=False,
            clear_search_url=this_url,
            search=view_args.search,

            # Filters
            filters=self._filters,
            filter_groups=self._get_filter_groups(),
            active_filters=view_args.filters,

            # Actions
            actions=actions,
            actions_confirmation=actions_confirmation,

            # Misc
            enumerate=enumerate,
            get_pk_value=self.get_pk_value,
            get_value=self.get_list_value,
            return_url=this_url,
        )

    def scaffold_sortable_columns(self):
        pass

    def get_pk_value(self, row):
        return row['key']

    @contextfunction
    def get_list_value(self, context, model, name):
        return model[name]

    def is_sortable(self, name):
        return False
